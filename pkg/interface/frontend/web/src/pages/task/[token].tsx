import { useParams } from 'react-router-dom'
import { Client, CommentArray, Equipment, EquipmentType, Order, User } from '../../models/api-models'
import { useEffect, useState } from 'react'
import axios from 'axios'
import Toolbar from '../../components/Toolbar'

interface ItemType {
    Order: Order
    Comments: CommentArray
    User: User
    Client: Client
    Equipment: Equipment
    EquipmentType: EquipmentType
}
function Index() {
    const params = useParams()
    const value = params?.token ?? ''
    const [item, setItem] = useState<ItemType>({} as ItemType)

    const GetData = async (value: string) => {
        try {
            const response = await axios.get('/api/Public/Tracking?Id=' + value)
            setItem(response.data)
        } catch (error) {
            console.log(error)
        }
    }
    useEffect(() => {
        GetData(value)
    }, [value])
    return (
        <div>
            <Toolbar />
            <OrderReceipt order={item?.Order} client={item?.Client} comments={item?.Comments} equipment={item?.Equipment} user={item?.User} />
        </div>
    )
}

interface OrderReceiptInfo {
    client?: Client
    order: Order
    equipment?: Equipment
    user?: User
    comments?: CommentArray
}

const OrderReceipt = ({ client, order, equipment, user, comments }: OrderReceiptInfo) => {
    return (
        <div style={{ fontFamily: 'Arial, sans-serif', margin: '20px' }}>
            <div className="container" style={{ width: '80%', margin: '0 auto' }}>
                {/* Header de la empresa */}
                <header style={{ textAlign: 'center', marginBottom: '30px' }}>
                    <img src="/logo.png" alt="Infoistel Logo" className="logo" style={{ width: '100px' }} />
                    <h1 style={{ margin: 0, color: '#333' }}>Infositel</h1>
                    <p style={{ margin: '5px 0', color: '#555' }}>Technology Solutions</p>
                    <p>
                        <strong>RUC:</strong> 123456789
                    </p>
                    <p>
                        <strong>Dirección:</strong> Calle Principal 123, Ciudad
                    </p>
                </header>

                {/* Información combinada en varias columnas */}
                <div className="section" style={{ marginBottom: '30px' }}>
                    <h2>Detalles de la Orden</h2>
                    <table style={{ width: '100%', borderCollapse: 'collapse', margin: '20px 0', border: '1px solid #CCC' }}>
                        <tbody style={{ border: '1px solid #CCC' }}>
                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>ID Cliente</th>
                                <td style={{ border: '1px solid #CCC' }}>{client?.ID ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>Nombre Cliente</th>
                                <td style={{ border: '1px solid #CCC' }}>{client?.Name ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>Ciudad</th>
                                <td style={{ border: '1px solid #CCC' }}>{client?.City?.String ?? ''}</td>
                            </tr>
                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>Dirección</th>
                                <td style={{ border: '1px solid #CCC' }}>{client?.Address?.String ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>Teléfono</th>
                                <td style={{ border: '1px solid #CCC' }}>{client?.Phone?.String ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>RUC</th>
                                <td style={{ border: '1px solid #CCC' }}>{client?.Ruc?.String ?? ''}</td>
                            </tr>

                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>Número de Orden</th>
                                <td style={{ border: '1px solid #CCC' }}>{order?.OrderNumber}</td>
                                <th style={{ border: '1px solid #CCC' }}>Fecha de Creación</th>
                                <td style={{ border: '1px solid #CCC' }}>{new Date(order?.CreationDate?.Time ?? '').toLocaleDateString()}</td>
                                <th style={{ border: '1px solid #CCC' }}>Prioridad</th>
                                <td style={{ border: '1px solid #CCC' }}>{order?.Priority?.String}</td>
                            </tr>
                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>Problema Reportado</th>
                                <td style={{ border: '1px solid #CCC' }} colSpan={5}>
                                    {order?.ReportedIssue?.String ?? ''}
                                </td>
                            </tr>
                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>Diagnóstico</th>
                                <td style={{ border: '1px solid #CCC' }} colSpan={5}>
                                    {order?.Diagnosis?.String ?? ''}
                                </td>
                            </tr>
                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>Solución Propuesta</th>
                                <td style={{ border: '1px solid #CCC' }} colSpan={5}>
                                    {order?.Solution?.String ?? ''}
                                </td>
                            </tr>

                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>ID Equipo</th>
                                <td style={{ border: '1px solid #CCC' }}>{equipment?.ID ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>Tipo de Equipo</th>
                                <td style={{ border: '1px solid #CCC' }}>{equipment?.TypeID?.Int32 ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>Serie del Equipo</th>
                                <td style={{ border: '1px solid #CCC' }}>{equipment?.SerialNumber ?? ''}</td>
                            </tr>

                            <tr style={{ border: '1px solid #CCC' }}>
                                <th style={{ border: '1px solid #CCC' }}>ID Usuario Asignado</th>
                                <td style={{ border: '1px solid #CCC' }}>{user?.ID ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>Nombre Usuario</th>
                                <td style={{ border: '1px solid #CCC' }}>{user?.Name ?? ''}</td>
                                <th style={{ border: '1px solid #CCC' }}>Email</th>
                                <td style={{ border: '1px solid #CCC' }}>{user?.Email ?? ''}</td>
                            </tr>
                        </tbody>
                    </table>
                </div>

                {/* Sección de Comentarios */}
                <div className="section" style={{ marginBottom: '30px' }}>
                    <h2>Comentarios</h2>
                    <table style={{ width: '100%', borderCollapse: 'collapse', margin: '20px 0' }}>
                        <thead>
                            <tr style={{ border: '1px solid #CCC' }}>
                                <th>Fecha</th>
                                <th>Comentario</th>
                            </tr>
                        </thead>
                        <tbody>
                            {comments?.map((comment, index) => (
                                <tr key={index}>
                                    <td style={{ border: '1px solid #CCC' }}>{new Date(comment?.Date?.Time).toLocaleDateString()}</td>
                                    <td style={{ border: '1px solid #CCC' }}>{comment?.Comment?.String ?? ''}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>

                {/* Footer */}
                <div className="footer" style={{ textAlign: 'center', marginTop: '50px', fontSize: '12px', color: '#888' }}>
                    <p>Infoistel - Technology Solutions</p>
                    <p>Teléfono: (123) 456-7890 | Email: contacto@infoistel.com</p>
                </div>
            </div>
        </div>
    )
}

export default Index
