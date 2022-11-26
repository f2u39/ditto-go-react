import * as React from 'react';
import { styled } from '@mui/material/styles';
import Card from '@mui/material/Card';
import CardMedia from '@mui/material/CardMedia';
import CardContent from '@mui/material/CardContent';
import CardActions from '@mui/material/CardActions';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import TuneIcon from '@mui/icons-material/Tune';
import PlayCircleOutlineIcon from '@mui/icons-material/PlayCircleOutline';
import { AppBar, Badge, Box, Button, Dialog, DialogActions, DialogContent, DialogTitle, Divider, FormControl, FormControlLabel, FormGroup, FormLabel, Grid, InputAdornment, InputLabel, MenuItem, Select, Switch, Tabs, TextField, ToggleButton, ToggleButtonGroup, Toolbar, Tooltip } from '@mui/material';
import Pagination from '@mui/material/Pagination';
import Tab from '@mui/material/Tab';
import TabContext from '@mui/lab/TabContext';
import TabList from '@mui/lab/TabList';
import TabPanel from '@mui/lab/TabPanel';
import { CheckSquareFill, Tablet, PcDisplay, NintendoSwitch, Playstation, Xbox } from 'react-bootstrap-icons';
import { Code, CodeSlash } from 'react-bootstrap-icons';
import { Battery, BatteryCharging, BatteryFull } from 'react-bootstrap-icons';
import { useEffect, useState } from 'react';
import PostAddIcon from '@mui/icons-material/PostAdd';
import AddBusinessIcon from '@mui/icons-material/AddBusiness';
import { useNavigate } from 'react-router-dom';

// import CreateIncDialog from './CreateIncDialog';

interface ExpandMoreProps extends IconButtonProps {
    expand: boolean
}

const ExpandMore = styled((props: ExpandMoreProps) => {
    const { expand, ...other } = props;
    return <IconButton {...other} />;
})(({ theme, expand }) => ({
    transform: !expand ? 'rotate(0deg)' : 'rotate(180deg)',
    marginLeft: 'auto',
    transition: theme.transitions.create('transform', {
        duration: theme.transitions.duration.shortest,
    }),
}));

export default function Game() {
    const navigate = useNavigate()
    const [details, setDetails] = useState<Detail[]>([])
    const [platform, setPlatform] = useState('All')
    const [page, setPage] = useState(1)
    const [totalPages, setTotalPages] = useState(1)
    const [status, setStatus] = useState("Playing")
    const [playedCount, setPlayedCount] = useState(0)
    const [playingCount, setPlayingCount] = useState(0)
    const [toPlayCount, setToPlayCount] = useState(0)
    const [openUpdateGameDialog, setOpenUpdateGameDialog] = useState(false)
    const [openCreateGameDialog, setOpenCreateGameDialog] = useState(false)

    const [openCreateIncDialog, setOpenCreateIncDialog] = useState(false)
    const handleCreateIncDialogOpen = () => { setOpenCreateIncDialog(true) }
    const handleCreateIncDialogClose = () => { setOpenCreateIncDialog(false) }

    const [createGame, setCreateGame] = useState({
        developers: [],
        publishers: [],
        genres: [],
        platforms: [],
    })

    const [updateGame, setUpdateGame] = useState({
        game: {
            id: String,
            title: String,
            genre: String,
            platform: String,
            developer_id: String,
            publisher_id: String,
            status: String,
            playtime: Number,
            rating: String,
        },
        developers: [],
        publishers: [],
        genres: [],
        platforms: [],
        play_time_hour: 0,
        play_time_min: 0
    })

    function fetchCreateGame() {
        fetch("/api/game/create", {
            method: "GET",
        })
            .then(resp => resp.json())
            .then(data => {
                if (data != null) {
                    setCreateGame(data)
                    setOpenCreateGameDialog(true)
                }
            })
    }

    function fetchUpdateGame(id: String) {
        fetch(`/api/game/update?id=${id}`, {
            method: "GET",
        })
            .then(resp => resp.json())
            .then(data => {
                if (data != null) {
                    setUpdateGame(data)
                    setOpenUpdateGameDialog(true)
                }
            })
    }

    const handleUpdateGameDialogOpen = (id: String) => {
        fetchUpdateGame(id)
    }
    const handleUpdateGameDialogClose = () => {
        setOpenUpdateGameDialog(false)
    }

    const handleCreateGameDialogOpen = () => {
        fetchCreateGame()
    }
    const handleCreateGameDialogClose = () => {
        setOpenCreateGameDialog(false)
    }

    function refresh() {
        fetch(`/api/game/status/${status}/${platform}/${page}`)
            .then(resp => resp.json())
            .then(data => {
                // console.log(data["details"])
                if (data["details"] != null) {
                    setDetails(data["details"])
                } else {
                    setDetails([])
                }

                setTotalPages(data["total_pages"])
            })

        fetch(`/api/game/counts`)
            .then(resp => resp.json())
            .then(data => {
                setPlayedCount(data["played_cnt"])
                setPlayingCount(data["playing_cnt"]) 
                setToPlayCount(data["toPlay_cnt"])
            })
    }

    useEffect(() => {
        refresh()
    }, [status, platform, page])

    const handleDeleteGame = (id: String) => {
        fetch(`/api/game/delete?id=${id}`, {
            method: "POST",
        })
        .then(() => {
            handleUpdateGameDialogClose()
            refresh()
        })
    }

    const handlePageChange = (event: React.ChangeEvent<unknown>, value: number) => {
        setPage(value)
    }

    const handleStatusChange = (event: React.SyntheticEvent, newStatus: string) => {
        setPage(1)
        setStatus(newStatus)
    }

    const handlePlatformChange = (event: React.SyntheticEvent, newValue: string) => {
        setPage(1)
        setPlatform(newValue)
    }

    const handleStartGame = (id: string) => {
        fetch(`/api/act/watch/start?id=${id}`)
    }

    return (
        <Box sx={{ width: '100%' }}>
            <Grid container sx={{ width: '50%' }}>
                <Grid item>
                    <Box justifyContent="flex-start">
                        <AppBar
                            sx={{ width: '50%', mr: '50%' }}
                            style={{ background: 'transparent', boxShadow: 'none' }}
                        >
                            <Toolbar>
                                <IconButton
                                    size="large"
                                    aria-controls="menu-appbar"
                                    aria-haspopup="true"
                                    color="inherit"
                                >
                                    <PostAddIcon sx={{ fontSize: 30, color: "thistle" }} onClick={handleCreateGameDialogOpen} />
                                </IconButton>

                                <IconButton
                                    size="large"
                                    aria-controls="menu-appbar"
                                    aria-haspopup="true"
                                    color="inherit"
                                >
                                    <AddBusinessIcon sx={{ fontSize: 30, color: "thistle" }} onClick={handleCreateIncDialogOpen} />
                                </IconButton>
                            </Toolbar>
                        </AppBar>
                    </Box>
                </Grid>
            </Grid>

            <TabContext value={status}>
                <Box
                    sx={{
                        borderBottom: 1,
                        borderColor: 'divider'
                    }}
                >
                    <TabList indicatorColor="secondary" onChange={handleStatusChange} centered>
                        <Tab
                            icon={
                                <Badge badgeContent={playedCount} color="primary">
                                    <BatteryFull fontSize="30" color="white" />
                                </Badge>
                            }
                            value="Played"
                        />
                        <Tab
                            icon={
                                <Badge badgeContent={playingCount} color="success">
                                    <BatteryCharging fontSize="30" color="green" />
                                </Badge>
                            }
                            value="Playing"
                        />
                        <Tab
                            icon={
                                <Badge badgeContent={toPlayCount} color="error">
                                    <Battery fontSize="30" color="red" />
                                </Badge>
                            }
                            value="ToPlay"
                        />
                    </TabList>
                </Box>

                <TabPanel value={status}>
                    <Grid
                        container
                        direction="row"
                        justifyContent="space-between"
                        sx={{ display: 'inline-flex' }}
                    >
                        <Grid
                            item
                            sx={{ m: -3, borderRight: 1, borderColor: 'divider' }}
                        >
                            <Tabs
                                variant="fullWidth"
                                orientation="vertical"
                                value={platform}
                                onChange={handlePlatformChange}
                            >
                                <Tab sx={{ mt: 6 }} icon={<CheckSquareFill size={30} />} value="All" />
                                <Tab sx={{ mt: 2 }} icon={<PcDisplay color="orange" size={30} />} value="PC" />
                                <Tab sx={{ mt: 2 }} icon={<Playstation color="#2E6DB4" size={30} />} value="PlayStation" />
                                <Tab sx={{ mt: 2 }} icon={<NintendoSwitch color="#E60012" size={30} />} value="Nintendo Switch" />
                                <Tab sx={{ mt: 2 }} icon={<Xbox color="#107C10" size={30} />} value="Xbox" />
                                <Tab sx={{ mt: 2 }} icon={<Tablet color="#730073" size={30} />} value="Mobile" />
                            </Tabs>
                        </Grid>

                        <Grid item xs={10}>
                            <Grid container>
                                {details.map((element, i) => (
                                    <Card
                                        sx={{ ml: 3, mt: 3, maxWidth: 250 }}
                                        key={element.game.id}
                                    >
                                        <CardMedia
                                            component="img"
                                            height="250"
                                            image={"/assets/images/games/" + element.game.id + ".webp"}
                                        />
                                        <CardContent>
                                            <Typography variant="subtitle1" align="center" color="text.secondary">
                                                {element.game.title}
                                            </Typography>
                                        </CardContent>

                                        <CardActions sx={{ mt: -1 }} disableSpacing>
                                            <Tooltip title="Property">
                                                <IconButton onClick={e => handleUpdateGameDialogOpen(element.game.id)}>
                                                    <TuneIcon />
                                                </IconButton>
                                            </Tooltip>

                                            <Tooltip title="Start">
                                                <IconButton onClick={e => handleStartGame(element.game.id)}>
                                                    <PlayCircleOutlineIcon />
                                                </IconButton>
                                            </Tooltip>
                                        </CardActions>

                                        <CardContent sx={{ mt: -4 }}>
                                            <Box sx={{
                                                mx: "auto",
                                                display: 'flex',
                                                flexDirection: 'column',
                                                alignItems: 'center',
                                                '& > *': {
                                                    m: 1,
                                                },
                                            }}
                                            >
                                                <TextField
                                                    fullWidth
                                                    size="small"
                                                    sx={{ pt: 1 }}
                                                    inputProps={{
                                                        style: { textAlign: 'right' },
                                                        readOnly: true,
                                                    }}
                                                    value={element.game.rating}
                                                    InputProps={{
                                                        startAdornment: (
                                                            <InputAdornment position="start">
                                                                {element.game.platform === 'Mobile' ? <Tablet /> : <></>}
                                                                {element.game.platform === 'PC' ? <PcDisplay /> : <></>}
                                                                {element.game.platform === 'Playstation' ? <Playstation /> : <></>}
                                                                {element.game.platform === 'Nintendo Switch' ? <NintendoSwitch /> : <></>}
                                                                {element.game.platform === 'Xbox' ? <Xbox /> : <></>}
                                                            </InputAdornment>
                                                        )
                                                    }}
                                                />

                                                <TextField
                                                    fullWidth
                                                    size="small"
                                                    sx={{ pt: 1 }}
                                                    inputProps={{
                                                        style: { textAlign: 'right' },
                                                        readOnly: true,
                                                    }}
                                                    value={element.developer.name}
                                                    InputProps={{
                                                        startAdornment: (
                                                            <InputAdornment position="start">
                                                                <Code />
                                                            </InputAdornment>
                                                        )
                                                    }}
                                                />

                                                <TextField
                                                    fullWidth
                                                    size="small"
                                                    sx={{ pt: 1 }}
                                                    inputProps={{ style: { textAlign: 'right' }, readOnly: true }}
                                                    value={element.publisher.name}
                                                    InputProps={{
                                                        startAdornment: (
                                                            <InputAdornment position="start"><CodeSlash /></InputAdornment>
                                                        )
                                                    }}
                                                />

                                                <TextField
                                                    fullWidth
                                                    size="small"
                                                    sx={{ pt: 1 }}
                                                    inputProps={{ style: { textAlign: 'right' }, readOnly: true }}
                                                    value={element.play_hour}
                                                    InputProps={{
                                                        startAdornment: (
                                                            <InputAdornment position="start">
                                                                {element.game.status === 'Played' ? <BatteryFull /> : <></>}
                                                                {element.game.status === 'Playing' ? <BatteryCharging /> : <></>}
                                                                {element.game.status === 'ToPlay' ? <Battery /> : <></>}
                                                            </InputAdornment>
                                                        ),
                                                        endAdornment: (
                                                            <InputAdornment position="end">Hour(s)</InputAdornment>
                                                        )
                                                    }}
                                                />
                                            </Box>
                                        </CardContent>
                                    </Card>
                                ))
                                }
                            </Grid>
                        </Grid>
                        <Grid xs={12} sx={{ pt: 3, pb: 3 }}>
                            <Box
                                display="flex"
                                justifyContent="center"
                                alignItems="center"
                            >
                                <Pagination
                                    count={totalPages}
                                    page={page}
                                    onChange={handlePageChange}
                                    variant="outlined"
                                    color="secondary" />
                            </Box>
                        </Grid>
                    </Grid>
                </TabPanel>
            </TabContext>

            <Dialog
                open={openUpdateGameDialog}
                onClose={handleUpdateGameDialogClose}
            >
                <DialogTitle align="center">
                    Update Game
                </DialogTitle>
                <DialogContent>
                    <form method="post" encType="multipart/form-data" action="/api/game/update">
                        <FormControl fullWidth sx={{ mt: 1 }}>
                            <TextField
                                name="id"
                                label="Id"
                                defaultValue={updateGame.game.id}
                                inputProps={{
                                    readOnly: true
                                }}
                            >
                            </TextField>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 2 }}>
                            <TextField
                                name="title"
                                label="Title"
                                defaultValue={updateGame.game.title}
                            >
                            </TextField>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 2 }}>
                            <InputLabel htmlFor="developer">Developer</InputLabel>
                            <Select
                                name="developer_id"
                                label="Developer"
                                defaultValue={updateGame.game.developer_id}
                            >
                                {updateGame.developers.map((dev: any, index) => {
                                    return (
                                        <MenuItem key={index} value={dev.id}>{dev.name}</MenuItem>
                                    )
                                })}
                            </Select>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 2 }}>
                            <InputLabel htmlFor="publisher">Publisher</InputLabel>
                            <Select
                                name="publisher_id"
                                label="Publisher"
                                defaultValue={updateGame.game.publisher_id}
                            >
                                {updateGame.publishers.map((pub: any, index) => {
                                    return (
                                        <MenuItem key={index} value={pub.id}>{pub.name}</MenuItem>
                                    )
                                })}
                            </Select>
                        </FormControl>

                        <Grid container>
                            <Grid item sx={{ width: '48%' }}>
                                <FormControl
                                    fullWidth
                                    sx={{ mt: 2 }}
                                >
                                    <InputLabel htmlFor="Status">Status</InputLabel>
                                    <Select
                                        name="status"
                                        label="Status"
                                        defaultValue={updateGame.game.status}
                                    >
                                        <MenuItem key="Played" value="Played">Played</MenuItem>
                                        <MenuItem key="Playing" value="Playing">Playing</MenuItem>
                                        <MenuItem key="ToPlay" value="ToPlay">ToPlay</MenuItem>
                                    </Select>
                                </FormControl>
                            </Grid>

                            <Grid item sx={{ width: '4%' }}></Grid>

                            <Grid item sx={{ width: '48%' }}>
                                <FormControl
                                    fullWidth
                                    sx={{ mt: 2 }}
                                >
                                    <InputLabel htmlFor="Genre">Genre</InputLabel>
                                    <Select
                                        name="genre"
                                        label="Genre"
                                        defaultValue={updateGame.game.genre}
                                    >
                                        {updateGame.genres.map((genre: any, index) => {
                                            return (
                                                <MenuItem key={index} value={genre}>{genre}</MenuItem>
                                            )
                                        })}
                                    </Select>
                                </FormControl>
                            </Grid>
                        </Grid>

                        <Grid container>
                            <Grid item sx={{ width: '48%' }}>
                                <FormControl
                                    fullWidth
                                    sx={{ mt: 2 }}
                                >
                                    <InputLabel htmlFor="Platform">Platform</InputLabel>
                                    <Select
                                        name="platform"
                                        label="Platform"
                                        defaultValue={updateGame.game.platform}
                                    >
                                        {updateGame.platforms.map((platform: any, index) => {
                                            return (
                                                <MenuItem key={index} value={platform}>{platform}</MenuItem>
                                            )
                                        })}
                                    </Select>
                                </FormControl>
                            </Grid>

                            <Grid item sx={{ width: '4%' }}></Grid>

                            <Grid item sx={{ width: '48%' }}>
                                <FormControl
                                    fullWidth
                                    sx={{ mt: 2 }}
                                >
                                    <InputLabel htmlFor="Rating">Rating</InputLabel>
                                    <Select
                                        name="rating"
                                        label="Rating"
                                        defaultValue={updateGame.game.rating}
                                    >
                                        <MenuItem key="S+" value="S+">S+</MenuItem>
                                        <MenuItem key="S" value="S">S</MenuItem>
                                        <MenuItem key="A+" value="A+">A+</MenuItem>
                                        <MenuItem key="A" value="A">A</MenuItem>
                                        <MenuItem key="B+" value="B+">B+</MenuItem>
                                        <MenuItem key="B" value="B">B</MenuItem>
                                        <MenuItem key="C+" value="C+">C+</MenuItem>
                                        <MenuItem key="C" value="C">C</MenuItem>
                                        <MenuItem key="D" value="D">D</MenuItem>
                                    </Select>
                                </FormControl>
                            </Grid>
                        </Grid>

                        <Grid container>
                            <Grid item sx={{ width: '23%' }}>
                                <FormControl fullWidth sx={{ mt: 2 }}>
                                    <TextField
                                        name="play_time_hour"
                                        type="number"
                                        label="Hour"
                                        defaultValue={updateGame.play_time_hour}
                                    >
                                    </TextField>
                                </FormControl>
                            </Grid>

                            <Grid item sx={{ width: '2%' }}></Grid>

                            <Grid item sx={{ width: '23%' }}>
                                <FormControl fullWidth sx={{ mt: 2 }}>
                                    <TextField
                                        name="play_time_min"
                                        type="number"
                                        label="Min"
                                        defaultValue={updateGame.play_time_min}
                                    >
                                    </TextField>
                                </FormControl>
                            </Grid>

                            <Grid item sx={{ width: '4%' }}></Grid>

                            <Grid item sx={{ width: '48%' }}>
                                <FormControl fullWidth sx={{ mt: 2 }}>
                                    <input type="file" id="cover" name="cover" />
                                </FormControl>
                            </Grid>
                        </Grid>
                        
                        <DialogActions style={{ justifyContent: "space-between" }} sx={{ mt: 1, mb: -1, ml: -1, mr: -1 }}>
                            <Button color="error" onClick={ e => handleDeleteGame(updateGame.game.id.toString()) }>Delete</Button>
                            <Box>
                                <Button color="secondary" onClick={handleUpdateGameDialogClose}>Cancel</Button>
                                <Button color="success" type="submit">Update</Button>
                            </Box>
                        </DialogActions>
                    </form>
                </DialogContent>
            </Dialog>

            <Dialog
                open={openCreateGameDialog}
                onClose={handleCreateGameDialogClose}
            >
                <DialogTitle align="center">Create Game</DialogTitle>
                <DialogContent>
                    <form method="post" action="/api/game/create">
                        <FormControl fullWidth sx={{ mt: 1 }}>
                            <TextField name="title" label="Title" required>
                            </TextField>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 2 }}>
                            <InputLabel htmlFor="developer">Developer</InputLabel>
                            <Select name="developer_id" label="Developer" required>
                                {createGame.developers.map((dev: any, index) => {
                                    return (<MenuItem key={index} value={dev.id}>{dev.name}</MenuItem>)
                                })}
                            </Select>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 2 }}>
                            <InputLabel htmlFor="publisher">Publisher</InputLabel>
                            <Select name="publisher_id" label="Publisher" required>
                                {createGame.publishers.map((pub: any, index) => {
                                    return (<MenuItem key={index} value={pub.id}>{pub.name}</MenuItem>)
                                })}
                            </Select>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 2 }}>
                            <InputLabel htmlFor="Genre">Genre</InputLabel>
                            <Select name="genre" label="Genre" required>
                                {createGame.genres.map((genre: any, index) => {
                                    return (<MenuItem key={index} value={genre}>{genre}</MenuItem>)
                                })}
                            </Select>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 2 }}>
                            <InputLabel htmlFor="Platform">Platform</InputLabel>
                            <Select name="platform" label="Platform" required>
                                {createGame.platforms.map((platform: any, index) => {
                                    return (
                                        <MenuItem key={index} value={platform}>{platform}</MenuItem>
                                    )
                                })}
                            </Select>
                        </FormControl>

                        <DialogActions sx={{ mt: 1, mb: -1, mr: -1 }}>
                            <Button color="secondary" onClick={handleCreateGameDialogClose}>Cancel</Button>
                            <Button color="success" type="submit">Create</Button>
                        </DialogActions>
                    </form>
                </DialogContent>
            </Dialog>

            <Dialog
                open={openCreateIncDialog}
                onClose={handleCreateIncDialogClose}
            >
                <DialogTitle align="center">Create Inc</DialogTitle>
                <DialogContent>
                    <form method="post" action="/api/inc/create">
                        <FormControl fullWidth sx={{ mt: 1, minWidth: 250 }}>
                            <TextField name="name" label="Name" required>
                            </TextField>
                        </FormControl>

                        <FormControl fullWidth sx={{ mt: 1 }} component="fieldset" variant="standard">
                            <FormGroup>
                                <FormControlLabel control={ <Switch value="1" name="is_developer" /> } label="Developer" />
                                <FormControlLabel control={ <Switch value="1" name="is_publisher" /> } label="Publisher" />
                            </FormGroup>
                        </FormControl>

                        <DialogActions sx={{ mt: 1, mb: -1, mr: -1 }}>
                            <Button color="secondary" onClick={handleCreateIncDialogClose}>Cancel</Button>
                            <Button color="success" type="submit">Create</Button>
                        </DialogActions>
                    </form>
                </DialogContent>
            </Dialog>
        </Box>
    )
}

interface Detail {
    game: Game,
    developer: Developer,
    publisher: Publisher,
    play_hour: 0,
    play_min: 0,
}

interface Game {
    id: string,
    title: string,
    genre: string,
    platform: string,
    developer_id: string,
    publisher_id: string,
    status: string,
    playtime: string,
    rating: string,
}

interface Developer {
    id: string,
    name: string,
}

interface Publisher {
    id: string,
    name: string,
}