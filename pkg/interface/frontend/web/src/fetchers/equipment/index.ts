import axios from 'axios'
import { CreateEquipmentQueryDTO, DeleteEquipmentTypeQueryDTO } from '../../models/api-models'

export const getEquipments = async () => {
    try {
        const response = await axios.get('/api/Equipment/List?offset=0&limit=200', {
            headers: {
                Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
            },
        })
        return response.data
    } catch (error) {
        throw new Error('Error fetching')
    }
}

export const getEquipementByID = async (id: number) => {
    try {
        const response = await axios.get('/api/Equipment/Get', {
            params: {
                id,
            },
            headers: {
                Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
            },
        })
        return response.data
    } catch (error) {
        throw new Error('Error fetching ')
    }
}

export const createEquipment = async (data: CreateEquipmentQueryDTO) => {
    try {
        const response = await axios.post('/api/Equipment/Create', data, {
            headers: {
                Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
            },
        })
        return response.data
    } catch (error) {
        throw new Error('Error fetching ')
    }
}

export const deleteEquipment = async (data: DeleteEquipmentTypeQueryDTO) => {
    try {
        const response = await axios.delete('/api/Equipment/Delete', {
            params: data,
            headers: {
                Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
            },
        })
        return response.data
    } catch (error) {
        throw new Error('Error fetching ')
    }
}
