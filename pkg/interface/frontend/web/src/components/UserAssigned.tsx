import { useQuery } from '@tanstack/react-query'
import { Avatar } from 'primereact/avatar'
import { getUserByID } from '../fetchers/user'
import { User } from '../models/api-models'
import { primeColors } from '../utils'

interface UserAssigned {
  IdUser: number
}

function UserAssigned({ IdUser }: UserAssigned) {
  const item = useQuery<User>({
    queryKey: ['userById-' + IdUser],
    refetchOnMount: true,
    queryFn: async () => await getUserByID(IdUser),
  })

  return (
    <div className="flex gap-2 align-items-center">
      <Avatar
        size="normal"
        shape="circle"
        label={item?.isLoading || item?.isError ? 'NN' : item.data?.Name.charAt(0)}
        style={{
          backgroundColor: primeColors[IdUser],
          color: 'white',
        }}
      />
      {item?.isLoading || item?.isError ? 'Cargando usuario...' : item.data?.Name}
    </div>
  )
}

export default UserAssigned
