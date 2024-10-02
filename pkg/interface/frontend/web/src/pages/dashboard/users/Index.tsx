import Toolbar from '../../../components/Toolbar'
import { DataTable } from 'primereact/datatable'
import { Column } from 'primereact/column'
import { UserArray } from '../../../models/api-models'
import { useQuery } from '@tanstack/react-query'
import { getUsers } from '../../../fetchers/user'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { Button } from 'primereact/button'

const home = {
  icon: 'pi pi-home',
  template: () => (
    <Link to="/dashboard">
      <i className={PrimeIcons.HOME}></i>
    </Link>
  ),
}

function Index() {
  const navigate = useNavigate()

  const items = [
    {
      label: 'Usuarios',
      command: () => navigate('/dashboard/users'),
    },
  ]
  const userQuery = useQuery<UserArray, Error>({
    queryKey: ['users'],
    queryFn: getUsers,
  })

  const header = (
    <div className="flex flex-wrap align-items-center justify-content-between gap-2">
      <span className="text-xl text-900 font-bold">Lista de usuarios</span>
      <Button
        icon={PrimeIcons.PLUS}
        label="Crear Nuevo usuario"
        onClick={() => navigate('/dashboard/users/create')}
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
          value={userQuery.data ?? []}
          header={header}
          loading={userQuery.isLoading}
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
          <Column field="Email" header="Email" sortable></Column>
          <Column field="PermissionType" header="Permission Type" sortable></Column>
          <Column field="CreationDate.Time" header="CreationDate " sortable></Column>
          <Column field="InactiveStatus.String" header="Status" sortable></Column>
          <Column field="Password.String" header="Password" sortable showFilterMenu filter></Column>
        </DataTable>
      </div>
    </div>
  )
}

export default Index
