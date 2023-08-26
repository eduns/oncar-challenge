import './index.css'

import { useEffect, useState } from 'react'

import Card from '../Card'

import { Vehicle } from '../../types'

export default function CardList() {
  const [vehicles, setVehicles] = useState<Vehicle[]>([])

  useEffect(() => {
    fetch('http://0.0.0.0:8080/vehicles')
      .then(response => response.json())
      .then(response => {
        setVehicles(response)
      })
      .catch(e => {
        console.error(e)
      })
  }, [])

  return (
    <>
      <h1>Escolha um veículo</h1>
      <div id="root">
        <div id='container'>
          {vehicles.map((vehicle: Vehicle, i) => (
            <Card
              key={i}
              vehicle={vehicle}
            />
          ))}
        </div>
      </div>
    </>
  )
}