import { useState } from 'react'
import Toolbar from '../../../components/Toolbar'
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { ClientArray, CreateRouteDTO, EquipmentArray, UserArray } from '../../../models/api-models'
import { createOrder } from '../../../fetchers/order'
import { getPanels } from '../../../fetchers/orderStatus'
import { getClients } from '../../../fetchers/client'
import { getUsers } from '../../../fetchers/user'
import { Dropdown } from 'primereact/dropdown'
import { Divider } from 'primereact/divider'
import { Button } from 'primereact/button'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { BreadCrumb } from 'primereact/breadcrumb'
import { InputNumber } from 'primereact/inputnumber'
import { InputTextarea } from 'primereact/inputtextarea'
import { getEquipments } from '../../../fetchers/equipment'

const home = {
    icon: 'pi pi-home',
    template: () => (
        <Link to="/dashboard">
            <i className={PrimeIcons.HOME}></i>
        </Link>
    ),
}

function Create() {
    const navigate = useNavigate()
    const queryClient = useQueryClient()
    const [values, setValues] = useState<CreateRouteDTO>({} as CreateRouteDTO)
    const createOrderMutation = useMutation({
        mutationFn: createOrder,
        onSuccess: async () => {
            queryClient.invalidateQueries({
                queryKey: ['orderstatus'],
            })
            queryClient.invalidateQueries({
                queryKey: ['clients'],
            })
            queryClient.invalidateQueries({
                queryKey: ['users'],
            })
            queryClient.invalidateQueries({
                queryKey: ['order'],
            })
        },
    })
    const statusQuery = useQuery<ClientArray, Error>({
        queryFn: getPanels,
        queryKey: ['orderstatus'],
    })
    const clientsQuery = useQuery<ClientArray, Error>({
        queryFn: getClients,
        queryKey: ['clients'],
    })
    const artifactsQuery = useQuery<EquipmentArray, Error>({
        queryFn: getEquipments,
        queryKey: ['artifacts'],
    })
    const userQuery = useQuery<UserArray, Error>({
        queryKey: ['users'],
        queryFn: getUsers,
    })

    const items = [
        {
            label: 'Ordernes',
            command: () => navigate('/dashboard/orders'),
        },
        {
            label: 'Crear Orden',
            command: () => navigate('/dashboard/orders/create'),
        },
    ]
    const generateRandomNumber = () => {
        const min = 100
        const max = 10000

        const randomFactor = Math.floor(Math.random() * (max - min + 1)) + min
        return Math.floor(randomFactor)
    }
    return (
        <>
            <Toolbar />
            <div className="m-3">
                <div className="mb-4">
                    <BreadCrumb model={items} home={home} />
                </div>
                <div className="grid">
                    <div className="xl:col-4 lg:col-4 md:col-4 sm:col-12">
                        <div>
                            <div className="mt-2">Nùmero de Order: </div>
                            <div className="mb-2 p-inputgroup flex-1">
                                <InputNumber prefix="Nª " className="w-full" disabled value={values.orderNumber as any} />
                                <Button
                                    icon={PrimeIcons.SYNC}
                                    className="p-button-primary"
                                    onClick={() => {
                                        setValues((prev) => ({ ...prev, orderNumber: generateRandomNumber().toString() }))
                                    }}
                                />
                            </div>
                        </div>
                        <div>
                            <div className="mt-2">Cliente: </div>
                            <div className="mb-2">
                                <Dropdown value={values.clientID} onChange={(e) => setValues((prev) => ({ ...prev, clientID: e.value }))} options={clientsQuery.data ?? []} optionLabel="Name" optionValue="ID" placeholder="Select a Client" className="mb-2 w-full" />
                            </div>
                        </div>
                        <div>
                            <div className="mt-2">Usuario asignado: </div>
                            <div className="mb-2">
                                <Dropdown value={values.assignedTo} onChange={(e) => setValues((prev) => ({ ...prev, assignedTo: e.value }))} options={userQuery.data ?? []} optionLabel="Name" optionValue="ID" placeholder="Select a User" className="mb-2 w-full" />
                            </div>
                        </div>
                        <Divider />
                        <div>
                            <div className="mt-2">Presupuesto: </div>
                            <div className="mb-2">
                                <InputNumber
                                    className="w-full"
                                    placeholder="S/. "
                                    onChange={(e) => {
                                        setValues((prev) => ({ ...prev, budget: e.value?.toString() ?? '' }))
                                    }}
                                />
                            </div>
                        </div>
                        <div>
                            <div className="mt-2">Tiempo estimado: </div>
                            <div className="mb-2">
                                <InputNumber
                                    prefix="&#9872; "
                                    suffix=" Horas"
                                    placeholder="&#9872; "
                                    min={1}
                                    max={40}
                                    className="w-full"
                                    value={values.estimatedTime as any}
                                    onChange={(e) => {
                                        setValues((prev) => ({ ...prev, estimatedTime: e.value as any }))
                                    }}
                                />
                            </div>
                        </div>
                    </div>
                    <div className="xl:col-8 lg:col-8 md:col-8 sm:col-12">
                        <div>
                            <div className="mt-2">Dispositivo: </div>
                            <div className="mb-2">
                                <Dropdown value={values.equipement} onChange={(e) => setValues((prev) => ({ ...prev, equipement: e.value }))} options={artifactsQuery.data ?? []} optionLabel="Name" optionValue="ID" placeholder="Selecciona el dispositivo" className="mb-2 w-full" />
                            </div>
                        </div>
                        <div>
                            {artifactsQuery?.data?.find((x) => x.ID === values.equipement) ? (
                                <>
                                    <div className="mt-2">Analisis previo: </div>
                                    <div className="mb-2 ">
                                        <InputTextarea disabled autoResize rows={3} className="w-full" value={artifactsQuery?.data?.find((x) => x.ID === values.equipement)?.Notes?.String ?? 'SIN ANÁLISIS PREVIO'} />
                                    </div>
                                </>
                            ) : (
                                <></>
                            )}
                            <div className="mt-2">Diagnòstico: </div>
                            <div className="mb-2">
                                <InputTextarea
                                    className="w-full"
                                    value={values.diagnosis}
                                    rows={5}
                                    onChange={(e) => {
                                        setValues((prev) => ({ ...prev, diagnosis: e.target.value }))
                                    }}
                                />
                            </div>
                        </div>
                        <div>
                            <div className="mt-2">Caràcteristica reportada: </div>
                            <div className="mb-2 p-inputgroup flex-1">
                                <InputTextarea
                                    rows={3}
                                    className="w-full"
                                    value={values.reportedIssue as any}
                                    onChange={(e) => {
                                        setValues((prev) => ({ ...prev, reportedIssue: e.target.value }))
                                    }}
                                />
                            </div>
                        </div>
                        <div>
                            <div className="mt-2">Soluciòn: </div>
                            <div className="mb-2 p-inputgroup flex-1">
                                <InputTextarea
                                    rows={3}
                                    className="w-full"
                                    value={values.solution as any}
                                    onChange={(e) => {
                                        setValues((prev) => ({ ...prev, solution: e.target.value }))
                                    }}
                                />
                            </div>
                        </div>
                    </div>
                </div>
                <div>
                    <Button
                        size="small"
                        onClick={async () => {
                            await createOrderMutation.mutateAsync({
                                id: values.id,
                                statusID: statusQuery?.data?.[0]?.ID ?? '',
                                assignedTo: values.assignedTo,
                                clientID: values.clientID,
                                priority: values.priority ?? 'medium',
                                budget: values.budget ?? '',
                                diagnosis: values.diagnosis ?? '',
                                equipement: values.equipement ?? 0,
                                estimatedTime: (((values.estimatedTime ?? 0) as any) * 3600) as any,
                                orderNumber: values.orderNumber ?? '',
                                reportedIssue: values.reportedIssue ?? '',
                                solution: values.solution ?? '',
                            } as CreateRouteDTO)
                        }}
                    >
                        Agregar Order
                    </Button>
                </div>
            </div>
        </>
    )
}

export default Create
