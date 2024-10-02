import React, { useState } from 'react'
import Toolbar from '../../components/Toolbar'
import { InputText } from 'primereact/inputtext'
import { Button } from 'primereact/button'
import { Message } from 'primereact/message'
import axios from 'axios'

interface FormData {
    correo: string
    asunto: string
    mensaje: string
}
function Index() {
    const [formData, setFormData] = useState<FormData>({
        correo: '',
        asunto: '',
        mensaje: '',
    })

    // Estado para el manejo de errores y loading
    const [errors, setErrors] = useState<Partial<FormData>>({})
    const [loading, setLoading] = useState<boolean>(false)
    const [formSent, setFormSent] = useState<boolean>(false)

    // Función para validar los inputs
    const validate = (): boolean => {
        const tempErrors: Partial<FormData> = {}

        if (!formData.correo || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.correo)) {
            tempErrors.correo = 'Por favor, ingrese un correo válido.'
        }

        if (!formData.asunto || formData.asunto.length < 5) {
            tempErrors.asunto = 'El asunto debe tener al menos 5 caracteres.'
        }

        if (!formData.mensaje || formData.mensaje.length < 10) {
            tempErrors.mensaje = 'El mensaje debe tener al menos 10 caracteres.'
        }

        setErrors(tempErrors)
        return Object.keys(tempErrors).length === 0
    }

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        setFormData({ ...formData, [e.target.name]: e.target.value })
        setErrors({ ...errors, [e.target.name]: '' })
    }

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (!validate()) return

        setLoading(true)
        try {
            await axios.post('/api/Public/SendSupportMessage', formData)
            setFormSent(true)
        } catch (error) {
            console.error('Error enviando el formulario:', error)
        } finally {
            setLoading(false)
        }
    }
    return (
        <div>
            <Toolbar />
            <div className="mt-2 px-3">Soporte Técnico</div>
            <div className="flex align-center justify-content-center p-mt-5 mt-5 px-3">
                <form className="p-fluid" onSubmit={handleSubmit} style={{ maxWidth: '500px' }}>
                    {/* Input de correo */}
                    <div className="p-field mb-2">
                        <label htmlFor="correo" className="mb-2">
                            Correo Electrónico
                        </label>
                        <InputText id="correo" name="correo" value={formData.correo} onChange={handleChange} className={errors.correo ? 'p-invalid' : ''} placeholder="Correo electrónico" />
                        <div className="my-2">{errors.correo && <Message severity="error" text={errors.correo} />}</div>
                    </div>

                    {/* Input de asunto */}
                    <div className="p-field mb-2">
                        <label htmlFor="asunto">Asunto</label>
                        <InputText id="asunto" name="asunto" value={formData.asunto} onChange={handleChange} className={errors.asunto ? 'p-invalid' : ''} placeholder="Asunto" />
                        <div className="my-2">{errors.asunto && <Message severity="error" text={errors.asunto} />}</div>
                    </div>

                    {/* Input de mensaje */}
                    <div className="p-field mb-2">
                        <label htmlFor="mensaje">Mensaje</label>
                        <textarea id="mensaje" name="mensaje" value={formData.mensaje} onChange={handleChange} rows={5} className={`p-inputtext p-component ${errors.mensaje ? 'p-invalid' : ''}`} placeholder="Describe tu problema" />
                        <div className="my-2">{errors.mensaje && <Message severity="error" text={errors.mensaje} />}</div>
                    </div>

                    {/* Botón de envío */}
                    <Button type="submit" label="Enviar" loading={loading} className="p-mt-2" />

                    {/* Mensaje de éxito */}
                    {formSent && (
                        <div className="mt-2">
                            <Message severity="success" text="¡Mensaje enviado exitosamente!" />
                        </div>
                    )}
                </form>
            </div>
        </div>
    )
}

export default Index
