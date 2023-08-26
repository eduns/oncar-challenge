import './index.css'

import { Link } from 'react-router-dom'

import { Vehicle } from '../../types'

export default function Card({ vehicle }: { vehicle: Vehicle}) {
  return (
    <div className='card'>
      <h3>{vehicle.model}</h3>
      <h4>{vehicle.year}</h4>
      <h5>{vehicle.brand}</h5>
      <h3 className='vehicle-price'>
        {new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(vehicle.price)}
      </h3>
      <Link to='/contact' className='choose-button' state={{chosenVehicleId: vehicle.id}}>Escolher veículo</Link>
    </div>
  )
}