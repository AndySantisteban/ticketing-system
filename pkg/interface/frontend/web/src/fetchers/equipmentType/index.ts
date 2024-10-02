import axios from 'axios'
import { CreateEquipmentTypeQueryDTO, DeleteEquipmentTypeQueryDTO } from '../../models/api-models'

export const getEquipmentsType = async () => {
  try {
    const response = await axios.get('/api/EquipmentType/List?offset=0&limit=200', {
      headers: {
        Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
      },
    })
    return response.data
  } catch (error) {
    throw new Error('Error fetching')
  }
}

export const getEquipementTypeByID = async (id: number) => {
  try {
    const response = await axios.get('/api/EquipmentType/Get', {
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

export const createEquipmentType = async (data: CreateEquipmentTypeQueryDTO) => {
  try {
    const response = await axios.post('/api/EquipmentType/Create', data, {
      headers: {
        Authorization: 'Bearer ' + localStorage.getItem('@token-infositel'),
      },
    })
    return response.data
  } catch (error) {
    throw new Error('Error fetching ')
  }
}

export const deleteEquipmentType = async (data: DeleteEquipmentTypeQueryDTO) => {
  try {
    const response = await axios.delete('/api/EquipmentType/Delete', {
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
