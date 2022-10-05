import { useState } from 'react';

export default function useToken() {
    const getToken = () => {
        const authToken = localStorage.getItem('auth_token');
        console.log(authToken);
        return authToken
    };

    const [token, setToken] = useState(getToken());

    const saveToken = (userToken: { auth_token: any; }) => {
        console.log(JSON.stringify(userToken.auth_token));
        localStorage.setItem('auth_token', JSON.stringify(userToken));
        setToken(userToken.auth_token);
        console.log(localStorage.getItem('auth_token'));
    };

    // return {
    //     token,
    //     setToken: saveToken
    // }

    return {
        token,
        setToken: saveToken
    }
}