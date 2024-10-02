import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query'
import { Button } from 'primereact/button'
import { Dialog } from 'primereact/dialog'
import { Divider } from 'primereact/divider'
import { Splitter, SplitterPanel } from 'primereact/splitter'
import { getOrdersByID, updateOrder } from '../fetchers/order'
import { CommentArray, Order, OrderStatusArray } from '../models/api-models'

import { PrimeIcons } from 'primereact/api'
import { Dropdown } from 'primereact/dropdown'
import { InputNumber } from 'primereact/inputnumber'
import { InputTextarea } from 'primereact/inputtextarea'
import { ScrollPanel } from 'primereact/scrollpanel'
import { useEffect, useState } from 'react'
import { createComment, getComment } from '../fetchers/comment'
import { getPanels } from '../fetchers/orderStatus'
import ClientAssigned, { ClientAssignedNotes } from './ClientAssigned'
import UserAssigned from './UserAssigned'

interface TaskProps {
    Id: number
    open: boolean
    onClose: () => void
}

function Task({ Id, open, onClose }: TaskProps) {
    const queryClient = useQueryClient()
    const [textComment, setTextComment] = useState<string>('')
    const [statusID, setStatusID] = useState<number>(0)
    const [amount, setAmount] = useState<string>('0')
    const item = useQuery<Order>({
        queryKey: ['orderByID-' + Id],
        queryFn: async () => await getOrdersByID(Id),
    })
    const listComment = useQuery<CommentArray>({
        queryKey: ['commentsByOrderID-' + Id],
        queryFn: async () => await getComment(Id),
    })

    const updateOrderMutation = useMutation({
        mutationKey: ['updateOrderByID-' + Id],
        mutationFn: updateOrder,
        onSuccess() {
            queryClient.invalidateQueries({
                queryKey: ['orderByID-' + Id],
            })
            queryClient.invalidateQueries({
                queryKey: ['activityByOrderID-' + Id],
            })
            queryClient.invalidateQueries({
                queryKey: ['commentsByOrderID-' + Id],
            })
            queryClient.invalidateQueries({
                queryKey: ['order'],
            })
            setAmount('0')
            setStatusID(1)
            onClose()
        },
    })

    const createCommentMutation = useMutation({
        mutationFn: createComment,
        mutationKey: ['createCommentBy-' + Id],
        onSuccess() {
            queryClient.invalidateQueries({
                queryKey: ['orderByID-' + Id],
            })
            queryClient.invalidateQueries({
                queryKey: ['activityByOrderID-' + Id],
            })
            queryClient.invalidateQueries({
                queryKey: ['commentsByOrderID-' + Id],
            })
        },
    })

    const status = useQuery<OrderStatusArray, Error>({
        queryKey: ['orderstatus'],
        refetchOnMount: true,
        queryFn: getPanels,
    })

    useEffect(() => {
        if (!item?.data?.ID) return
        setStatusID(item?.data.StatusID.Int32)
        setAmount(item?.data?.Budget?.String)
    }, [item.data])

    if (item?.isLoading || item?.isError) {
        return <Dialog visible={open} onHide={onClose} style={{ width: '70%', height: '90%' }}></Dialog>
    }

    return (
        <Dialog
            visible={open}
            onHide={onClose}
            header={'Orden Nª ' + item?.data?.OrderNumber}
            footer={
                <Button
                    label={'Guardar cambios'}
                    size="small"
                    icon={PrimeIcons.CHECK}
                    onClick={() => {
                        updateOrderMutation.mutateAsync({
                            assignedTo: item?.data?.AssignedTo?.Int32 ?? 0,
                            budget: item?.data?.Budget?.String ?? '',
                            clientID: item?.data?.ClientID?.Int32 ?? 0,
                            creationDate: item?.data?.CreationDate?.Time ?? new Date(),
                            diagnosis: item?.data?.Diagnosis?.String ?? '',
                            equipement: item?.data?.EquipmentID?.Int32 ?? 0,
                            estimatedTime: item?.data?.EstimatedTime?.Int64 as BigInt,
                            id: item?.data?.ID ?? 0,
                            orderNumber: item?.data?.OrderNumber as string,
                            priority: item?.data?.Priority?.String ?? '',
                            reportedIssue: item?.data?.ReportedIssue?.String ?? '',
                            solution: item?.data?.Solution?.String ?? '',
                            statusID: item?.data?.StatusID?.Int32 ?? 0,
                        })
                    }}
                />
            }
        >
            <ScrollPanel style={{ width: '100%', height: '350px' }}>
                <Splitter style={{ height: '100%', borderColor: 'transparent' }}>
                    <SplitterPanel size={75} minSize={10}>
                        <div className="px-1 w-full">
                            {item?.data?.ID && <ClientAssignedNotes Id={item?.data?.ID} />}
                            <Divider />
                            <div>
                                <div className="flex justify-content-between align-items-center gap-3">
                                    <div style={{ width: '90%' }}>
                                        <InputTextarea placeholder="Inserte un comentario" value={textComment} onChange={(e) => setTextComment(e.target.value ?? '')} style={{ height: '50px' }} className="w-full" />
                                    </div>
                                    <div>
                                        <Button
                                            icon={PrimeIcons.CHECK}
                                            label="Guardar"
                                            severity="warning"
                                            outlined
                                            size="small"
                                            onClick={() => {
                                                createCommentMutation.mutateAsync({
                                                    Comment: textComment ?? 'No insertò el comentario',
                                                    Date: new Date(),
                                                    ID: 0,
                                                    OrderID: Id,
                                                    UserID: 6,
                                                })
                                                setTextComment('')
                                            }}
                                        />
                                    </div>
                                </div>
                                <h4 className="">Comentarios recientes</h4>
                                <ScrollPanel style={{ width: '100%', height: '170px' }}>
                                    <div>
                                        {listComment?.data?.reverse()?.map((x) => (
                                            <div>
                                                <UserAssigned IdUser={x?.UserID?.Int32} />
                                                <small className="mx-5">{x.Comment?.String ?? ''}</small>
                                            </div>
                                        ))}
                                    </div>
                                </ScrollPanel>
                            </div>
                        </div>
                    </SplitterPanel>
                    <SplitterPanel size={25}>
                        <div className="px-3 w-full">
                            <div className="w-full">
                                <Dropdown
                                    value={statusID}
                                    onChange={(e) => {
                                        setStatusID(e.value)
                                    }}
                                    options={status.data ?? []}
                                    optionLabel="Name"
                                    optionValue="ID"
                                    className="w-full"
                                />
                            </div>
                            <p>User Assigned:</p>
                            <>
                                {item?.data?.AssignedTo && <UserAssigned IdUser={item?.data?.AssignedTo?.Int32} />}
                                <p>Client:</p>
                                {item?.data?.ClientID && <ClientAssigned Id={item?.data?.ClientID?.Int32} />}
                            </>
                            <div className="w-full">
                                <p>Amount :</p>
                                <InputNumber inputId="currency-pe" value={Number(amount ?? '0')} onValueChange={(e) => setAmount(e.value?.toString() ?? '0')} mode="currency" currency="PEN" locale="es-PE" />
                            </div>
                        </div>
                    </SplitterPanel>
                </Splitter>
            </ScrollPanel>
        </Dialog>
    )
}

export default Task
