import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import 'primeflex/primeflex.css' // flex
import 'primeicons/primeicons.css' //icons
import { ProgressSpinner } from 'primereact/progressspinner'
import 'primereact/resources/primereact.min.css' //core css
import 'primereact/resources/themes/lara-light-indigo/theme.css' //theme
// import 'primereact/resources/themes/mdc-light-indigo/theme.css'
import { Suspense } from 'react'
import { BrowserRouter, useRoutes } from 'react-router-dom'
import routes from '~react-pages'
import './App.css'
function Routes() {
  const element = useRoutes(routes)
  return element
}
const queryClient = new QueryClient()

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Suspense
          fallback={
            <div
              style={{
                minHeight: '100vh',
                minWidth: '100%',
                display: 'flex',
                alignItems: 'center',
                justifyContent: 'center',
              }}
            >
              <ProgressSpinner strokeWidth="8" />
            </div>
          }
        >
          <Routes />
        </Suspense>
      </BrowserRouter>
    </QueryClientProvider>
  )
}

export default App
