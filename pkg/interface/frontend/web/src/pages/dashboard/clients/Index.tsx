import React from 'react'
import Toolbar from '../../../components/Toolbar'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { useQuery } from '@tanstack/react-query'
import { ClientArray } from '../../../models/api-models'
import { getClients } from '../../../fetchers/client'
import { Button } from 'primereact/button'
import { BreadCrumb } from 'primereact/breadcrumb'
import { DataTable } from 'primereact/datatable'
import { Column } from 'primereact/column'

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
            label: 'Clientes',
            command: () => navigate('/dashboard/clients'),
        },
    ]
    const clientQuery = useQuery<ClientArray, Error>({
        queryKey: ['clients'],
        queryFn: getClients,
    })

    const header = (
        <div className="flex flex-wrap align-items-center justify-content-between gap-2">
            <span className="text-xl text-900 font-bold">Lista de Clientes</span>
            <Button icon={PrimeIcons.PLUS} label="Crear Nuevo Cliente" size="small" onClick={() => navigate('/dashboard/clients/create')} />
        </div>
    )
    return (
        <div>
            <Toolbar />
            <div className="m-3">
                <div className="mb-4">
                    <BreadCrumb model={items} home={home} />
                </div>
                <DataTable value={clientQuery.data ?? []} header={header} loading={clientQuery.isLoading} showGridlines stripedRows paginator sortMode="multiple" removableSort rows={5} rowsPerPageOptions={[5, 10, 25, 50]}>
                    <Column field="ID" header="ID" sortable resizeable></Column>
                    <Column field="Name" header="Name" sortable resizeable></Column>
                    <Column field="Address.String" header="Address" sortable resizeable></Column>
                    <Column field="District.String" header="District" sortable></Column>
                    <Column field="City.String" header="City" sortable></Column>
                    <Column field="Country.String" header="Country" sortable></Column>
                    <Column field="Phone.String" header="Phone" sortable></Column>
                    <Column field="Ruc.String" header="RUC" sortable></Column>
                    <Column field="ContactPerson.String" header="Contact Person" sortable></Column>
                    <Column field="Email.String" header="Email" sortable></Column>
                    <Column field="Website.String" header="Website" sortable></Column>
                    <Column field="AddressLine2.String" header="Address Line 2" sortable></Column>
                    <Column field="PostalCode.String" header="Postal Code" sortable></Column>
                    <Column field="Fax.String" header="Fax" sortable></Column>
                    <Column field="Notes.String" header="Notes" sortable></Column>
                </DataTable>
            </div>
        </div>
    )
}

export default Index
