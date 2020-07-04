import React from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import TopBar from '../../components/appBar/topBar'
import Room from '../room'

export default function App() {
  return (
    <div>
      <CssBaseline />
      <TopBar />
      <Room />
    </div>
  )
}
