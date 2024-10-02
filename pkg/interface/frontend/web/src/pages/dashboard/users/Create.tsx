import { Button } from 'primereact/button'
import Toolbar from '../../../components/Toolbar'
import { InputText } from 'primereact/inputtext'
import { useMutation, useQueryClient } from '@tanstack/react-query'
import { useState } from 'react'
import { CreateUserDTO } from '../../../models/api-models'
import { createUsers } from '../../../fetchers/user'
import { Dropdown } from 'primereact/dropdown'
import { Password } from 'primereact/password'
import { BreadCrumb } from 'primereact/breadcrumb'
import { Link, useNavigate } from 'react-router-dom'
import { PrimeIcons } from 'primereact/api'

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
  const [values, setValues] = useState<CreateUserDTO>({} as CreateUserDTO)
  const createuserMutation = useMutation({
    mutationFn: createUsers,
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
      label: 'Usuarios',
      command: () => navigate('/dashboard/users'),
    },
    {
      label: 'Crear Nuevo Usuario',
      command: () => navigate('/dashboard/users/create'),
    },
  ]

  return (
    <>
      <Toolbar />

      <div className="mx-3 pt-3">
        <div className="mb-4">
          <BreadCrumb model={items} home={home} />
        </div>
        <div className="flex">
          <div className="w-full">
            <div>
              <div className="mt-2">Nombre: </div>
              <div className="mb-2">
                <InputText
                  value={values.name ?? ''}
                  onChange={(e) => setValues((prev) => ({ ...prev, name: e.target.value }))}
                  placeholder="Inserta un nombre de usuario"
                  className="mb-2 w-full"
                />
              </div>
            </div>
            <div>
              <div className="mt-2">Correo electronico: </div>
              <div className="mb-2">
                <InputText
                  value={values.email ?? ''}
                  onChange={(e) => setValues((prev) => ({ ...prev, email: e.target.value ?? '' }))}
                  placeholder="user@infositel.com"
                  // mask="a*a*a*a*a@infositel.com"
                  className="mb-2 w-full"
                />
              </div>
            </div>
            <div>
              <div className="mt-2">Tipo de usuario: </div>
              <div className="mb-2 w-full">
                <Dropdown
                  options={[
                    { value: 'admin', label: 'Administrador' },
                    { value: 'supervisor', label: 'Supervisor' },
                    { value: 'employee', label: 'Colaborador' },
                  ]}
                  value={values.permission_type}
                  optionLabel="label"
                  optionValue="value"
                  onChange={(e) =>
                    setValues((prev) => ({
                      ...prev,
                      permission_type: e.value ?? '',
                    }))
                  }
                  className="mb-2 w-full"
                />
              </div>
            </div>
            <div>
              <div className="mt-2">Contraseña: </div>
              <div className="mb-2 w-full">
                <Password
                  value={values.password ?? ''}
                  onChange={(e) =>
                    setValues((prev) => ({
                      ...prev,
                      password: e.target.value ?? '',
                    }))
                  }
                  id="password"
                  toggleMask
                  className="mb-2 w-full"
                />
              </div>
            </div>
          </div>
        </div>
        <div className="flex align-items-center justify-content-end w-100">
          <Button
            onClick={async () => {
              await createuserMutation.mutateAsync({
                id: 0,
                creation_date: new Date(),
                email: values.email ?? '',
                inactive_status: 'active',
                name: values.name ?? '',
                password: values.password ?? '',
                permission_type: values.permission_type ?? '',
              } as CreateUserDTO as any)
            }}
          >
            Añadir nuevo usuario
          </Button>
        </div>
      </div>
    </>
  )
}

export default Create
