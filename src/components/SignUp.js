import React,{useState} from "react";
import axios from "axios";

const signup = () => {
    const [userName,setUsername] = useState('');
    const [password,setPassword] = useState('');
    
    const handleSignUp = (e) => {
        e.preventDefault();
        axios.post('http://localhost:9000/signup')
    }

}