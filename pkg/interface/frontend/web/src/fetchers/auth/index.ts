import axios from 'axios'

export const getPermission = async () => {
    try {
        const response = await axios.get('/api/login/userType', {
            headers: {
                Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
            },
        })
        if (response.data === '<nil>') {
            localStorage.removeItem('@token-infositel')
            // window.location.href = "/";
            return
        }
        return response.data
    } catch (error) {
        throw new Error('Error fetching users')
    }
}

export const getUserName = async () => {
    try {
        const response = await axios.get('/api/login/userName', {
            headers: {
                Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
            },
        })
        if (response.data === '<nil>') {
            localStorage.removeItem('@token-infositel')
            // window.location.href = '/'
            return
        }
        return response.data
    } catch (error) {
        throw new Error('Error fetching users')
    }
}
