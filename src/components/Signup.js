import React,{useState} from "react";
import axios from "axios";
import './styling/signup.css'


const Signup = () => {
    const [formData,setFormData] = useState({
    first_name: '',
    last_name : '',
    password : '',
    email : '',
    phone : '',
    user_type : 'USER'
    });
    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try{
            const response = await axios.post('http://localhost:9000/users/signup',formData)
            console.log("user signed up successfully:",response.data)
            window.location.href = "/login"
        }
        catch(error){
            console.error("error signing up user:",error.response ? error.response.data : error.message);
        }
    };
    return(
        <form onSubmit={handleSubmit} className="form-container">
            <h2>Sign Up</h2>
            <div>
                <label>First Name :</label>
                <input type="text" name="first_name" value={formData.first_name} onChange={handleChange}></input>
            </div>
            <div>
                <label>Last Name :</label>
                <input type="text" name="last_name" value={formData.last_name} onChange={handleChange}></input>
            </div>
            <div>
                <label>Email :</label>
                <input type="text" name="email" value={formData.email} onChange={handleChange}></input>
            </div>
            <div>
                <label>Password :</label>
                <input type="text" name="password" value={formData.password} onChange={handleChange}></input>
            </div>
            <div>
                <label>Mobile Number :</label>
                <input type="text" name="phone" value={formData.phone} onChange={handleChange}></input>
            </div>
            <button type="submit">Sign Up</button>
        </form>
    )

}

export default Signup