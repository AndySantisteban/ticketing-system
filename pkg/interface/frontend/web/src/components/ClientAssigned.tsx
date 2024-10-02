import { useQuery } from '@tanstack/react-query'
import { Avatar } from 'primereact/avatar'
import { Editor } from 'primereact/editor'
import { getClientByID } from '../fetchers/client'
import { Client, Order } from '../models/api-models'
import { primeColors } from '../utils'
import { getOrdersByID } from '../fetchers/order'

interface ClienAssigned {
    Id: number
}
function ClientAssigned({ Id }: ClienAssigned) {
    const item = useQuery<Client>({
        queryKey: ['clienById-' + Id],
        queryFn: async () => await getClientByID(Id),
    })

    return (
        <div className="flex gap-2 align-items-center">
            <Avatar
                size="normal"
                shape="circle"
                label={item?.isLoading || item?.isError ? 'NN' : item.data?.Name.charAt(0)}
                style={{
                    backgroundColor: primeColors[Id],
                    color: 'white',
                }}
            />
            {item?.isLoading || item?.isError ? 'Loading User...' : item.data?.Name}
        </div>
    )
}

export function ClientAssignedNotes({ Id }: ClienAssigned) {
    const item = useQuery<Order>({
        queryKey: ['orderByIdNotes-' + Id],
        queryFn: async () => await getOrdersByID(Id),
    })
    return (
        <div className="w-full">
            <Editor disabled value={item.data?.Diagnosis?.String ?? ''} style={{ height: '120px' }} className="w-full" />
        </div>
    )
}

export default ClientAssigned
