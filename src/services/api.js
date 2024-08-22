import axios from "axios";

const API_URL = 'http://localhost:9000'


const getAuthHeaers = () => {
    const token = localStorage.getItem('jwtToken');
    return {
        headers: {
            'token' : `${token}`
        },
    };
};


export const getLeaveList = async () => {
    try {
        const token = localStorage.getItem('token');
        console.log(`${token}`)
        const response = await axios.get(`${API_URL}/leave_list`, {
            headers: {
                
                'token': `${token}`
            }
        });
        return response.data;
    } catch (error) {
        throw error;
    }
};


