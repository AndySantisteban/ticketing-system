import { useState } from 'react'
import Toolbar from '../../../components/Toolbar'
import { CreateEquipmentQueryDTO, EquipmentTypeArray } from '../../../models/api-models'
import { Link, useNavigate } from 'react-router-dom'
import { useMutation, useQuery } from '@tanstack/react-query'
import { createEquipment } from '../../../fetchers/equipment'
import { BreadCrumb } from 'primereact/breadcrumb'
import { PrimeIcons } from 'primereact/api'
import { Divider } from 'primereact/divider'
import { Button } from 'primereact/button'
import { InputText } from 'primereact/inputtext'
import { InputTextarea } from 'primereact/inputtextarea'
import { Dropdown } from 'primereact/dropdown'
import { getEquipmentsType } from '../../../fetchers/equipmentType'

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
  const [inputs, setInputs] = useState<CreateEquipmentQueryDTO>({} as CreateEquipmentQueryDTO)
  const items = [
    {
      label: 'Equipos',
      command: () => navigate('/dashboard/artifacts'),
    },
    {
      label: 'Crear Equipo',
      command: () => navigate('/dashboard/artifacts/create'),
    },
  ]
  const action = useMutation({
    mutationFn: createEquipment,
    onSuccess: () => {
      navigate('/dashboard/artifacts')
    },
  })
  const qArtifacts = useQuery<EquipmentTypeArray, Error>({
    queryKey: ['typeArtifacts'],
    queryFn: getEquipmentsType,
  })

  return (
    <div>
      <Toolbar />
      <div className="m-3">
        <div className="mb-4">
          <BreadCrumb model={items} home={home} />
        </div>
        <div className="grid">
          <div className="xl:col-6 lg:col-6 md:col-6 sm:col-12">
            <div>
              <p>Nombre:</p>
              <InputText
                className="w-full"
                value={inputs?.name}
                onChange={(e) => setInputs((prev) => ({ ...prev, name: e.target.value }))}
              />
            </div>
            <div>
              <p>Notas:</p>
              <InputTextarea
                className="w-full"
                value={inputs?.notes}
                onChange={(e) => setInputs((prev) => ({ ...prev, notes: e.target.value }))}
              />
            </div>
            <Divider />
          </div>
          <div className="xl:col-6 lg:col-6 md:col-6 sm:col-12">
            <div>
              <p>N° de Serie:</p>
              <InputText
                className="w-full"
                value={inputs?.serial_number}
                onChange={(e) => setInputs((prev) => ({ ...prev, serial_number: e.target.value }))}
              />
            </div>
            <div>
              <p>Tipo de equipo:</p>
              <Dropdown
                value={inputs?.type_id}
                onChange={(e) => setInputs((prev) => ({ ...prev, type_id: e.value }))}
                options={qArtifacts.data}
                optionLabel="Name"
                optionValue="ID"
                placeholder="Selecciona un tipo de Equipo"
                className="w-full"
              />
            </div>
          </div>
          <Divider />
        </div>
        <div>
          <Button
            label="crear"
            size="small"
            onClick={() => {
              action.mutateAsync({
                id: 0,
                name: inputs?.name,
                notes: inputs?.notes ?? '',
                serial_number: inputs?.serial_number ?? '',
                type_id: inputs?.type_id ?? 0,
              })
            }}
          ></Button>
        </div>
      </div>
    </div>
  )
}

export default Create
