import React, {useState} from "react";
import axios from "axios";
import Cookies  from 'universal-cookie';
import {useHistory} from "react-router-dom";

const Login = () => {
    const cookies = new Cookies();
    const history = useHistory();
    const [username,setUsername] = useState('');
    const [password,setPassword] = useState('');

    const onSubmitLogin = (e : React.SyntheticEvent) => {
        e.preventDefault();
        axios.post('http:/login', {
            username,
            password
        }).then(res => {
            const token : TokenInterface = res.data;
            cookies.set('token', token.token, { path: '/' });
            const path = '/'
            history.push(path);
        }).catch(err => console.log(err))
    }

    return <div className="login-wrapper">

        <form >
            <h2>Welcome</h2>
            <div className="input-group">
                <input type="text" id="username" title="jgjjg" required value={username} onChange={e => setUsername(e.target.value)}/>
                <label htmlFor="username">Username</label>
            </div>
            <div className="input-group">
                <input type="password" id="password" required value={password} onChange={e => setPassword(e.target.value) }/>
                <label htmlFor="password">Password</label>
            </div>
            {/*<div className="input-group">*/}
            {/*    <input type="checkbox"/>*/}
            {/*    <label htmlFor="">Remember me</label>*/}
            {/*</div>*/}
            <button type="submit" onClick={onSubmitLogin}>Login</button>
        </form>


    </div>
}

export default Login;