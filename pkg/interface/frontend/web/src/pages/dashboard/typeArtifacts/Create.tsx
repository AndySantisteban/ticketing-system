import { useState } from 'react'
import Toolbar from '../../../components/Toolbar'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'
import { useMutation } from '@tanstack/react-query'
import { createEquipmentType } from '../../../fetchers/equipmentType'
import { InputText } from 'primereact/inputtext'
import { CreateEquipmentTypeQueryDTO } from '../../../models/api-models'
import { Button } from 'primereact/button'

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
  const [inputs, setInputs] = useState<CreateEquipmentTypeQueryDTO>(
    {} as CreateEquipmentTypeQueryDTO
  )

  const items = [
    {
      label: 'Tipos de Equipos',
      command: () => navigate('/dashboard/typeArtifacts'),
    },
    {
      label: 'Crear Tipo de Equipo',
      command: () => navigate('/dashboard/typeArtifacts/create'),
    },
  ]
  const createAction = useMutation({
    mutationFn: createEquipmentType,
    onSuccess: () => {
      navigate('/dashboard/typeArtifacts')
    },
  })
  return (
    <div>
      <Toolbar />
      <div className="m-3">
        <div className="mb-4">
          <BreadCrumb model={items} home={home} />
        </div>
        <div className="mb-2">
          <p>Inserta un nuevo tipo de Equipo:</p>
          <InputText
            value={inputs.name}
            onChange={(e) => setInputs((prev) => ({ ...prev, name: e.target.value }))}
          />
        </div>
        <div>
          <Button
            label="Crear"
            onClick={() => {
              createAction.mutateAsync({
                id: 0,
                name: inputs?.name ?? '',
              })
            }}
          />
        </div>
      </div>
    </div>
  )
}

export default Create
