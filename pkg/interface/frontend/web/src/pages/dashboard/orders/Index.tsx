import Toolbar from '../../../components/Toolbar'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { useQuery } from '@tanstack/react-query'
import { Order, OrderArray } from '../../../models/api-models'
import { Button } from 'primereact/button'
import { BreadCrumb } from 'primereact/breadcrumb'
import { DataTable } from 'primereact/datatable'
import { Column } from 'primereact/column'
import { getOrders } from '../../../fetchers/order'

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
            label: 'Ordenes',
            command: () => navigate('/dashboard/orders'),
        },
    ]
    const orderQuery = useQuery<OrderArray, Error>({
        queryKey: ['orders'],
        queryFn: getOrders,
    })

    const header = (
        <div className="flex flex-wrap align-items-center justify-content-between gap-2">
            <span className="text-xl text-900 font-bold">Lista de Ordenes</span>
            <Button icon={PrimeIcons.PLUS} label="Crear Nueva Orden" onClick={() => navigate('/dashboard/orders/create')} />
        </div>
    )
    return (
        <div>
            <Toolbar />
            <div className="m-3">
                <div className="mb-4">
                    <BreadCrumb model={items} home={home} />
                </div>
                <DataTable value={orderQuery.data ?? []} header={header} loading={orderQuery.isLoading} showGridlines stripedRows paginator sortMode="multiple" removableSort rows={5} rowsPerPageOptions={[5, 10, 25, 50]}>
                    <Column field="ID" header="ID" />
                    <Column field="ClientID.Int32" header="Client ID" body={(rowData: Order) => rowData.ClientID.Int32} />
                    <Column field="EquipmentID.Int32" header="Equipment ID" body={(rowData: Order) => rowData.EquipmentID.Int32} />
                    <Column field="OrderNumber" header="Order Number" />
                    <Column field="ReportedIssue.String" header="Reported Issue" body={(rowData: Order) => rowData.ReportedIssue.String} />
                    <Column field="Diagnosis.String" header="Diagnosis" body={(rowData: Order) => rowData.Diagnosis.String} />
                    <Column field="Solution.String" header="Solution" body={(rowData: Order) => rowData.Solution.String} />
                    <Column field="EstimatedTime.Int64" header="Estimated Time" body={(rowData: Order) => rowData.EstimatedTime.Int64.toString()} />
                    <Column field="Budget.String" header="Budget" body={(rowData: Order) => rowData.Budget.String} />
                    <Column field="StatusID.Int32" header="Status ID" body={(rowData: Order) => rowData.StatusID.Int32} />
                    <Column field="AssignedTo.Int32" header="Assigned To" body={(rowData: Order) => rowData.AssignedTo.Int32} />
                    <Column field="CreationDate.Time" header="Creation Date" body={(rowData: Order) => rowData.CreationDate.Time.toLocaleString()} />
                    <Column field="Priority.String" header="Priority" body={(rowData: Order) => rowData.Priority.String} />
                </DataTable>
            </div>
        </div>
    )
}

export default Index
