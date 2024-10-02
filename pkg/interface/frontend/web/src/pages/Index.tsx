import axios from 'axios'
import { Button } from 'primereact/button'
import { InputText } from 'primereact/inputtext'
import { Password } from 'primereact/password'
import { Toast } from 'primereact/toast'
import { useEffect, useRef, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Logo from '/logo.png'
function Index() {
  const [username, setUsername] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [token, setToken] = useState<string>(localStorage.getItem('@token-infositel') ?? '')
  const navigate = useNavigate()
  const toast = useRef<Toast>(null)

  const handleLogin = async () => {
    try {
      const response = await axios.post('/api/login/auth', {
        username: username,
        password: password,
      })
      toast?.current?.show?.({
        severity: 'info',
        summary: 'Auth',
        detail: 'User Logged',
      })
      setTimeout(() => {
        localStorage.setItem('@token-infositel', response.data)
        setToken(response.data)
      }, 2000)
    } catch (e) {
      toast?.current?.show?.({
        severity: 'error',
        summary: 'Auth',
        detail: 'Credentials not valid',
      })
    }
  }

  useEffect(() => {
    if (token) {
      navigate('/dashboard')
    }
  }, [token, navigate])
  return (
    <>
      <Toast ref={toast} />

      <div className="h-screen flex w-screen flex align-items-center justify-content-center">
        <div className="flex align-items-center justify-content-center ">
          <div className="p-shadow-3 mb-3 px-3">
            <img src={Logo} alt="Infositel" />
            <h1 className="text-start mb-1">Login</h1>
            <h2 className="text-start text-orange-500 mt-0">Infositel</h2>

            <div className="p-field">
              <p className="my-1">Username:</p>
              <InputText
                id="username"
                className="w-full"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
              />
            </div>
            <div className="p-field mb-2">
              <p className="my-1">Password:</p>
              <Password
                id="password"
                value={password}
                toggleMask
                inputClassName="w-full"
                className="bg-red-100 w-full"
                pt={{ iconField: { root: { className: 'w-full' } } }}
                feedback={false}
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
            <Button
              label="Login"
              icon="pi pi-user"
              onClick={handleLogin}
              size="small"
              severity="warning"
              className="w-full"
            />
          </div>
        </div>
      </div>
    </>
  )
}
export default Index
