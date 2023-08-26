import { SyntheticEvent } from 'react'
import './index.css'

import { useLocation } from 'react-router-dom'

export default function FormLead() {
  const { state } = useLocation()
  const { chosenVehicleId } = state
  
  const checkFieldErrors = (msg: string): string[] => {
    const errors = msg.split('\n')
    const fieldsWithErrors:string[] = []

    if (errors.find(e => e.includes("Error:Field validation for 'Name'"))) {
      fieldsWithErrors.push('Nome')
    }

    if (errors.find(e => e.includes("Error:Field validation for 'Email'"))) {
      fieldsWithErrors.push('Email')
    }
    
    if (errors.find(e => e.includes("Error:Field validation for 'Phone'"))) {
      fieldsWithErrors.push('Telefone')
    }

    return fieldsWithErrors
  }

  function handleSubmit (event: SyntheticEvent) {
    event.preventDefault()

    const { name, email, phone, chosenVehicleId } = event.target as typeof event.target & {
      name: { value: string };
      email: { value: string };
      phone: { value: string };
      chosenVehicleId: { value: string };
    }

    const msgDiv = document.querySelector('div#errors') as HTMLDivElement
    const form = document.querySelector('form#leadInfo') as HTMLFormElement
    
    msgDiv.style.visibility = 'hidden'

    fetch('http://0.0.0.0:8080/leads', {
      method: 'POST',
      body: JSON.stringify({
        'name': name.value,
        'email': email.value,
        'phone': '+' + phone.value,
        'chosenVehicleId': chosenVehicleId.value
      })
    })
      .then(response => {
        if (response.status !== 201) {
          return response.json()
        }
      })
      .then(response => {
        msgDiv.style.visibility = 'visible'

        if (response && response['error']) {
          const fieldsWithErrors = checkFieldErrors(response['error'])

          if (fieldsWithErrors.length > 0) {
            msgDiv.innerText = 'Erros nos campos: ' + fieldsWithErrors.join(', ')
            return
          }

          msgDiv.innerText = 'Erro ao enviar informações'

          console.error(response['error'])
          return
        }

        form.reset()
        msgDiv.innerText = 'Tudo certo! ;D'
      })
      .catch(e => {
        console.error(e)
      })
  }

  return (
    <>
      <h3>Preencha as informações de contato</h3>
      <div className='container'>
        <form id="leadInfo" onSubmit={handleSubmit}>
          <div>
            <input type='text' name='name' placeholder='Nome' required />
          </div>
          <div>
            <input type='email' name='email' placeholder='E-mail' required />
          </div>
          <div>
            <input type='tel' name='phone' placeholder='Telefone (Ex: 5512984761375)' pattern='55[0-9]{2}[0-9]{9}' required />
          </div>
          <input type='hidden' name='chosenVehicleId' value={chosenVehicleId} style={{'visibility': 'hidden'}}/>
          <button type='submit'>Enviar</button>
        </form>
        <div id='errors'></div>
      </div>
    </>
  )
}