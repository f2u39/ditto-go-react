package firebase

import (
	"bytes"
	"context"
	"ditto/model/todo"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"reflect"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const IS_DEP_ENV = false // Deployment environment flag

var (
	bkt *storage.BucketHandle
	cli *firestore.Client
	ctx context.Context
)

func All(col string, result interface{}) error {
	iter := cli.Collection(col).Documents(ctx)
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("Result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()
	i := 0
	for {
		elemp := reflect.New(elemt)
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Failed to iterate:", err)
			return err
		}
		if err := doc.DataTo(elemp.Interface()); err != nil {
			log.Println("Failed to convert:", err)
			return err
		}

		slicev = reflect.Append(slicev, elemp.Elem())
		slicev = slicev.Slice(0, slicev.Cap())
		i++
	}
	resultv.Elem().Set(slicev.Slice(0, i))
	return nil
}

func ById(col string, id string, T interface{}) error {
	dsnap, err := cli.Collection(col).Doc(id).Get(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return dsnap.DataTo(&T)
}

func Create(c string, m map[string]interface{}) error {
	_, _, err := cli.Collection(c).Add(ctx, m)
	return err
}

func DeleteId(col string, id string) error {
	_, err := cli.Collection(col).Doc(id).Delete(ctx)
	return err
}

func Delete(col string, path string, op string, value interface{}) error {
	iter := cli.Collection(col).Where(path, op, value).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		var todo todo.Todo
		if err := doc.DataTo(&todo); err != nil {
			return err
		}

		err = DeleteId(col, todo.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func Update(col string, id string, m map[string]interface{}) error {
	_, err := cli.Collection(col).Doc(id).Set(ctx, m, firestore.MergeAll)

	if err != nil {
		log.Println("An error has occurred:", err)
	}
	return err
}

func Init() {
	var err error
	var cfg *firebase.Config
	var opt option.ClientOption
	ctx = context.Background()

	if IS_DEP_ENV {
		cfg = &firebase.Config{
			StorageBucket: "ditto-375b1.appspot.com",
		}
		opt = option.WithCredentialsFile("config/ditto-375b1-firebase-adminsdk-5f9e4-392e525037.json")
	} else {
		cfg = &firebase.Config{
			StorageBucket: "ditto-7d7ad.appspot.com",
		}
		opt = option.WithCredentialsFile("config/ditto-7d7ad-f4d0196febeb.json")
	}

	app, err := firebase.NewApp(ctx, cfg, opt)
	if err != nil {
		log.Println(err)
	}

	// Storage client
	stoCli, err := app.Storage(context.Background())
	if err != nil {
		log.Println(err)
	}

	bkt, err = stoCli.DefaultBucket()
	if err != nil {
		log.Println(err)
	}

	// Firestore client
	cli, err = app.Firestore(ctx)
	if err != nil {
		log.Println(err)
	}
	// defer cli.Close()
}

// func Upload(fileInput []byte, fileName string) error
func Upload(header *multipart.FileHeader, rename string) error {
	// TODO The correct way to use context?
	// ctx, cancel := context.WithTimeout(context.Background(), fs.defaultTransferTimeout)
	// defer cancel()

	if bkt == nil {
		return fmt.Errorf("bucket is not initialized")
	}

	file, err := header.Open()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	object := bkt.Object(rename)
	writer := object.NewWriter(ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, bytes.NewReader(data)); err != nil {
		return err
	}

	if err := object.ACL().Set(context.Background(), storage.AllUsers, storage.RoleReader); err != nil {
		return err
	}

	return nil
}
