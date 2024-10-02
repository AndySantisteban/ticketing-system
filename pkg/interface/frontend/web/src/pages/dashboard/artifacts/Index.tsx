import Toolbar from '../../../components/Toolbar'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { useMutation, useQuery } from '@tanstack/react-query'
import { ClientArray, Equipment } from '../../../models/api-models'
import { Button } from 'primereact/button'
import { BreadCrumb } from 'primereact/breadcrumb'
import { DataTable } from 'primereact/datatable'
import { Column } from 'primereact/column'
import { deleteEquipment, getEquipments } from '../../../fetchers/equipment'
import { Dispatch, SetStateAction, useState } from 'react'
import { Dialog } from 'primereact/dialog'

const home = {
    icon: 'pi pi-home',
    template: () => (
        <Link to="/dashboard">
            <i className={PrimeIcons.HOME}></i>
        </Link>
    ),
}

const DeleteItem = ({ item, open, setOpen }: { item: Equipment; open: boolean; setOpen: Dispatch<SetStateAction<boolean>> }) => {
    const action = useMutation({
        mutationKey: ['deleteArtifacts'],
        mutationFn: deleteEquipment,
        onSettled: () => {
            setOpen(!open)
        },
    })
    return (
        <Dialog
            visible={open}
            onHide={() => setOpen(!open)}
            footer={() => (
                <div>
                    <Button
                        label="Eliminar"
                        onClick={() => {
                            action.mutateAsync({
                                id: item?.ID,
                            })
                        }}
                    />
                </div>
            )}
        >
            ¿Seguro que deseas eliminar esto?
        </Dialog>
    )
}

function Index() {
    const navigate = useNavigate()
    const [item, setItem] = useState<Equipment>({} as Equipment)
    const [open, setOpen] = useState<boolean>(false)

    const items = [
        {
            label: 'Equipos',
            command: () => navigate('/dashboard/artifacts'),
        },
    ]
    const query = useQuery<ClientArray, Error>({
        queryKey: ['artifacts'],
        queryFn: getEquipments,
    })

    const header = (
        <div className="flex flex-wrap align-items-center justify-content-between gap-2">
            <span className="text-xl text-900 font-bold">Lista de equipos</span>
            <Button icon={PrimeIcons.PLUS} label="Crear Nuevo Equipos" onClick={() => navigate('/dashboard/artifacts/create')} />
        </div>
    )
    return (
        <div>
            <Toolbar />
            <div className="m-3">
                <div className="mb-4">
                    <BreadCrumb model={items} home={home} />
                </div>
                <DataTable value={query.data ?? []} header={header} loading={query.isLoading} showGridlines stripedRows paginator sortMode="multiple" removableSort rows={5} rowsPerPageOptions={[5, 10, 25, 50]}>
                    <Column field="Name" header="Name" sortable></Column>
                    <Column
                        body={(data) => {
                            return (
                                <Button
                                    size="small"
                                    severity="danger"
                                    label="Eliminar"
                                    onClick={() => {
                                        setOpen(true)
                                        setItem(data)
                                    }}
                                />
                            )
                        }}
                    ></Column>
                </DataTable>
            </div>
            <DeleteItem item={item} open={open} setOpen={setOpen} />
        </div>
    )
}

export default Index
