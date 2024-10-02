import { useState } from 'react'
import Toolbar from '../../../components/Toolbar'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { CreateClientRouteDTO } from '../../../models/api-models'
import { createClient } from '../../../fetchers/client'
import { Divider } from 'primereact/divider'
import { Button } from 'primereact/button'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { BreadCrumb } from 'primereact/breadcrumb'
import { InputText } from 'primereact/inputtext'
import { InputTextarea } from 'primereact/inputtextarea'

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
    const [values, setValues] = useState<CreateClientRouteDTO>({} as CreateClientRouteDTO)
    const createMutation = useMutation({
        mutationFn: createClient,
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

    const items = [
        {
            label: 'Clientes',
            command: () => navigate('/dashboard/clients'),
        },
        {
            label: 'Crear Cliente',
            command: () => navigate('/dashboard/clients/create'),
        },
    ]
    return (
        <>
            <Toolbar />
            <div className="m-3">
                <div className="mb-4">
                    <BreadCrumb model={items} home={home} />
                </div>
                <div className="flex">
                    <div style={{ width: '50%' }}>
                        <div className="mt-2">Nombre: (*) </div>
                        <div className="mb-2">
                            <InputText value={values.name ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, name: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Direccion: (*)</div>
                        <div className="mb-2">
                            <InputText value={values.address ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, address: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Contacto para aviso: (*)</div>
                        <div className="mb-2">
                            <InputText value={values.contactPerson ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, contactPerson: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Ciudad: (*)</div>
                        <div className="mb-2">
                            <InputText value={values.city ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, city: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Distrito: (*)</div>
                        <div className="mb-2">
                            <InputText value={values.district ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, district: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Codigo postal: (*)</div>
                        <div className="mb-2">
                            <InputText value={values.postalCode ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, postalCode: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                    </div>
                    <Divider layout="vertical" />
                    <div style={{ width: '50%' }}>
                        <div className="mt-2">Correo electronico: (*)</div>
                        <div className="mb-2">
                            <InputText value={values.email ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, email: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Fax: </div>
                        <div className="mb-2">
                            <InputText value={values.fax ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, fax: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Direccion secundaria: </div>
                        <div className="mb-2">
                            <InputText
                                value={values.addressLine2 ?? ''}
                                onChange={(e) =>
                                    setValues((prev) => ({
                                        ...prev,
                                        addressLine2: e.target.value,
                                    }))
                                }
                                placeholder=""
                                className="mb-2 w-full"
                            />
                        </div>
                        <div className="mt-2">RUC/DNI: (*) </div>
                        <div className="mb-2">
                            <InputText value={values.ruc ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, ruc: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Telefono: (*) </div>
                        <div className="mb-2">
                            <InputText value={values.phone ?? ''} onChange={(e) => setValues((prev) => ({ ...prev, phone: e.target.value }))} placeholder="" className="mb-2 w-full" />
                        </div>
                        <div className="mt-2">Nota: </div>
                        <div className="mb-2">
                            <InputTextarea
                                value={values.notes ?? ''}
                                onChange={(e) =>
                                    setValues((prev) => ({
                                        ...prev,
                                        notes: e.target.value,
                                    }))
                                }
                                placeholder=""
                                className="mb-2 w-full form-select"
                            />
                        </div>
                    </div>
                </div>
                <div className="flex align-items-center justify-content-end w-100">
                    <Button
                        onClick={async () => {
                            await createMutation.mutateAsync({
                                address: values.address ?? '',
                                addressLine2: values.addressLine2 ?? '',
                                city: values.city ?? '',
                                contactPerson: values.contactPerson ?? '',
                                country: values.country ?? 'PE',
                                district: values.district ?? '',
                                email: values.email ?? '',
                                fax: values.fax ?? '',
                                name: values.name ?? '',
                                notes: values.notes ?? '',
                                phone: values.phone ?? '',
                                id: 0,
                                postalCode: values.postalCode ?? '',
                                ruc: values.ruc ?? '',
                                website: values.website ?? '',
                            } as CreateClientRouteDTO)
                        }}
                    >
                        Añadir nuevo Cliente
                    </Button>
                </div>
            </div>
        </>
    )
}

export default Create
