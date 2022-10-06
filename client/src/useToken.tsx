import { useState } from 'react';

export default function useToken() {
    const getToken = () => {
        // const authToken = localStorage.getItem('auth_token');
        const authToken = getCookie("auth_token")
        return authToken
    };

    const [token, setToken] = useState(getToken());
    const saveToken = (userToken: { auth_token: any; }) => {
        // localStorage.setItem('auth_token', JSON.stringify(userToken));
        setCookie("auth_token", JSON.stringify(userToken))
        setToken(userToken.auth_token)
    };

    return {
        token,
        setToken: saveToken
    }

    function setCookie(name: string, val: string) {
        // const expire = new Date();
        // expire.setTime(expire.getTime() + (8 * 60 * 60 * 1000));
        // document.cookie = name+"="+val+"; expires="+expire.toUTCString()+"; path=/";
        document.cookie = name+"="+val+"; expires=; path=/";
    }

    function getCookie(name: string): string {
	return document.cookie
		.split(';')
		.map(c => c.trim())
		.filter(cookie => {
			return cookie.substring(0, name.length + 1) === `${name}=`;
		})
		.map(cookie => {
			return decodeURIComponent(cookie.substring(name.length + 1));
		})[0];
    }
}