import Toolbar from '../../../components/Toolbar'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { useMutation, useQuery } from '@tanstack/react-query'
import { EquipmentType, EquipmentTypeArray } from '../../../models/api-models'
import { Button } from 'primereact/button'
import { BreadCrumb } from 'primereact/breadcrumb'
import { DataTable } from 'primereact/datatable'
import { Column } from 'primereact/column'
import { deleteEquipmentType, getEquipmentsType } from '../../../fetchers/equipmentType'
import { Dialog } from 'primereact/dialog'
import { Dispatch, SetStateAction, useState } from 'react'

const home = {
  icon: 'pi pi-home',
  template: () => (
    <Link to="/dashboard">
      <i className={PrimeIcons.HOME}></i>
    </Link>
  ),
}
const DeleteItem = ({
  item,
  open,
  setOpen,
}: {
  item: EquipmentType
  open: boolean
  setOpen: Dispatch<SetStateAction<boolean>>
}) => {
  const action = useMutation({
    mutationKey: ['deleteTypeArtifacts'],
    mutationFn: deleteEquipmentType,
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
                Id: item?.ID,
              })
            }}
          />
        </div>
      )}
    >
      {JSON.stringify(item)}
    </Dialog>
  )
}
function Index() {
  const navigate = useNavigate()
  const [open, setOpen] = useState<boolean>(false)
  const [item, setItem] = useState<EquipmentType>({} as EquipmentType)

  const items = [
    {
      label: 'Tipos de Equipos',
      command: () => navigate('/dashboard/typeArtifacts'),
    },
  ]
  const query = useQuery<EquipmentTypeArray, Error>({
    queryKey: ['typeArtifacts'],
    queryFn: getEquipmentsType,
  })

  const header = (
    <div className="flex flex-wrap align-items-center justify-content-between gap-2">
      <span className="text-xl text-900 font-bold">Lista de Tipos de equipos</span>
      <Button
        icon={PrimeIcons.PLUS}
        label="Crear Nuevo Tipo de Equipos"
        onClick={() => navigate('/dashboard/typeArtifacts/create')}
      />
    </div>
  )
  return (
    <div>
      <Toolbar />
      <div className="m-3">
        <div className="mb-4">
          <BreadCrumb model={items} home={home} />
        </div>
        <DataTable
          value={query?.data ?? []}
          header={header}
          loading={query.isLoading}
          showGridlines
          stripedRows
          paginator
          sortMode="multiple"
          removableSort
          rows={5}
          rowsPerPageOptions={[5, 10, 25, 50]}
        >
          <Column field="ID" header="ID" sortable></Column>
          <Column field="Name" header="Name" sortable></Column>
          <Column
            body={(data) => {
              return (
                <Button
                  label="Eliminar"
                  severity="danger"
                  onClick={() => {
                    setItem(data)
                    setOpen(true)
                  }}
                />
              )
            }}
          />
        </DataTable>
      </div>
      <DeleteItem item={item} setOpen={setOpen} open={open} />
    </div>
  )
}

export default Index
