import { useQuery } from '@tanstack/react-query'
import { PrimeIcons } from 'primereact/api'
import { Avatar } from 'primereact/avatar'
import { AvatarGroup } from 'primereact/avatargroup'
import { Button } from 'primereact/button'
import { Card } from 'primereact/card'
import { Chip } from 'primereact/chip'
import { Panel } from 'primereact/panel'
import { ScrollPanel } from 'primereact/scrollpanel'
import { Tooltip } from 'primereact/tooltip'
import { Fragment, useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { getOrders } from '../fetchers/order'
import { getPanels } from '../fetchers/orderStatus'
import { getUsers } from '../fetchers/user'
import { Order, OrderArray, OrderStatusArray, UserArray } from '../models/api-models'
import Task from './Task'
import { primeColors } from '../utils'

function KanbanBoard() {
    const [itemselected, setItemSelected] = useState<Order>({} as Order)
    const [showDetailItem, setDetailItem] = useState<boolean>(false)
    const [userSelected, setUserSelected] = useState<string>('')
    const navigate = useNavigate()
    const panelsQuery = useQuery<OrderStatusArray, Error>({
        queryKey: ['orderstatus'],
        refetchOnMount: true,
        queryFn: getPanels,
    })
    const ordersQUery = useQuery<OrderArray, Error>({
        queryKey: ['order'],
        refetchOnMount: true,
        queryFn: getOrders,
    })
    const userQuery = useQuery<UserArray, Error>({
        queryKey: ['users'],
        refetchOnMount: true,
        queryFn: getUsers,
    })

    const AvatarsUser = useMemo(() => {
        return (
            <>
                <AvatarGroup>
                    <Avatar
                        onClick={() => setUserSelected('')}
                        title={'+'}
                        size="large"
                        shape="circle"
                        style={{
                            backgroundColor: primeColors[18],
                            color: 'white',
                        }}
                    />
                    {userQuery.data?.map((x) => {
                        return (
                            <Fragment key={x.ID}>
                                <Tooltip target={'.custom-target-icon-' + x.ID} />
                                <Avatar
                                    key={x.ID}
                                    onClick={() => {
                                        setUserSelected(x.ID.toString())
                                    }}
                                    title={x.Name}
                                    label={x.Name.charAt(0)}
                                    className={'custom-target-icon-' + x.ID}
                                    size="large"
                                    shape="circle"
                                    style={{
                                        backgroundColor: primeColors[x.ID],
                                        color: 'white',
                                    }}
                                />
                            </Fragment>
                        )
                    })}
                </AvatarGroup>
            </>
        )
    }, [userQuery.data])

    return (
        <div className="mt-4  px-5">
            {itemselected?.ID && <Task Id={itemselected.ID} onClose={() => setDetailItem(false)} open={showDetailItem} />}
            <div>
                <h2>Infositel</h2>
            </div>
            <div className="mb-3">{AvatarsUser}</div>
            <div className="flex gap-2">
                <div className="w-full flex gap-2 overflow-x-scroll">
                    {panelsQuery.data?.map((x) => {
                        return (
                            <Panel header={x.Name ?? ''} key={x.ID} style={{ width: '350px', minWidth: '350px' }} className="pb-3">
                                <ScrollPanel style={{ width: '100%', height: '450px' }}>
                                    {ordersQUery?.data
                                        ?.filter((o) => o.StatusID.Int32 === x.ID)
                                        .filter((x) => (userSelected === '' ? true : Number(userSelected) === x.AssignedTo.Int32))
                                        ?.map((x) => {
                                            return (
                                                <Card
                                                    onClick={() => {
                                                        setItemSelected(x)
                                                        setDetailItem(true)
                                                    }}
                                                    style={{ cursor: 'pointer' }}
                                                    key={x.ID}
                                                    className="my-2 border-solid border-300 border-1"
                                                    subTitle={x.Diagnosis?.String ?? 'Without Title'}
                                                    title={x.OrderNumber}
                                                >
                                                    <div className="flex align-items-center justify-content-end gap-2">
                                                        <Chip label={x.Priority?.String} icon={PrimeIcons.ANGLE_DOUBLE_UP + ' text-primary'} />
                                                        <Button
                                                            icon={PrimeIcons.EXTERNAL_LINK}
                                                            size="small"
                                                            onClick={() => {
                                                                navigate('/task/' + x.ID)
                                                            }}
                                                        />
                                                    </div>
                                                </Card>
                                            )
                                        })}
                                </ScrollPanel>
                            </Panel>
                        )
                    })}
                </div>
            </div>
        </div>
    )
}

export default KanbanBoard
