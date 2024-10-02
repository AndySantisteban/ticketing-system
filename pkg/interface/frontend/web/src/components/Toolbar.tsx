import { useQuery } from '@tanstack/react-query'
import { PrimeIcons } from 'primereact/api'
import { Button } from 'primereact/button'
import { Menubar } from 'primereact/menubar'
import { useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { getPermission, getUserName } from '../fetchers/auth'

function usePermisions() {
    const navigate = useNavigate()
    const [openCreate, setOpenCreate] = useState<boolean>(false)
    const [openCreateUser, setOpenCreateUser] = useState<boolean>(false)
    const [openCreateClient, setOpenCreateClient] = useState<boolean>(false)
    const auth = useQuery<string>({
        queryKey: ['userType'],
        queryFn: async () => await getPermission(),
        refetchOnMount: true,
    })
    const userName = useQuery<string>({
        queryKey: ['userName'],
        queryFn: async () => await getUserName(),
        refetchOnMount: true,
    })

    const items = useMemo(() => {
        const itemsUserAdmin = [
            {
                label: 'Panel Kanban',
                icon: PrimeIcons.TH_LARGE,
                command: () => navigate('/dashboard'),
            },
            {
                label: 'Clientes',
                icon: 'pi pi-users',
                command: () => navigate('/dashboard/clients'),
            },
            {
                label: 'Usuarios',
                icon: 'pi pi-user',
                command: () => navigate('/dashboard/users'),
            },

            {
                label: 'Ordernes',
                icon: 'pi pi-inbox',
                command: () => navigate('/dashboard/orders'),
            },
            {
                label: 'Equipos',
                icon: PrimeIcons.TAG,
                items: [
                    {
                        label: 'Tipos de equipos',
                        icon: PrimeIcons.BOOKMARK,
                        command: () => navigate('/dashboard/typeArtifacts'),
                    },
                    {
                        label: 'Todos los equipos registrados',
                        icon: PrimeIcons.BOX,
                        command: () => navigate('/dashboard/Artifacts'),
                    },
                ],
            },
        ]
        const itemsUserSupervisor = [
            {
                label: 'Panel Kanban',
                icon: PrimeIcons.TH_LARGE,
                command: () => navigate('/dashboard'),
            },
            {
                label: 'Clientes',
                icon: 'pi pi-users',
                command: () => navigate('/dashboard/clients'),
            },
            {
                label: 'Ordernes',
                icon: 'pi pi-inbox',
                command: () => navigate('/dashboard/orders'),
            },
            {
                label: 'Tipos de artefactos',
                icon: PrimeIcons.TAG,
                command: () => navigate('/dashboard/artifacts'),
            },
        ]
        const itemsUserEmployee = [
            {
                label: 'Panel Kanban',
                icon: PrimeIcons.TH_LARGE,
                command: () => navigate('/dashboard'),
            },
            {
                label: 'Ordernes',
                icon: 'pi pi-inbox',
                command: () => navigate('/dashboard/orders'),
            },
            {
                label: 'Tipos de artefactos',
                icon: PrimeIcons.TAG,
                command: () => navigate('/dashboard/artifacts'),
            },
        ]
        const list = auth.data === 'supervisor' ? itemsUserSupervisor : auth.data === 'admin' ? itemsUserAdmin : auth.data === undefined ? [] : itemsUserEmployee
        return [
            ...list,
            {
                label: 'Soporte Tecnico',
                icon: 'pi pi-envelope',
                command: () => navigate('/support'),
            },
        ]
    }, [auth?.data, navigate])
    return {
        userType: auth.data,
        items,
        order: { openCreate, setOpenCreate },
        user: { openCreateUser, setOpenCreateUser },
        client: { openCreateClient, setOpenCreateClient },
        userName: userName.data,
    }
}
function Toolbar() {
    const permission = usePermisions()
    const navigate = useNavigate()
    const EndContent = () => (
        <div className="flex align-items-center gap-1">
            <div style={{ fontWeight: 'bold', textTransform: 'uppercase' }}>{permission.userName ?? ''}</div>
            <div>
                <Button
                    severity="warning"
                    text
                    icon={PrimeIcons.SIGN_OUT}
                    onClick={() => {
                        localStorage?.removeItem('@token-infositel')
                        navigate('/')
                    }}
                    size="small"
                />
            </div>
        </div>
    )

    return <Menubar model={permission.items ?? []} end={EndContent} />
}

export default Toolbar
