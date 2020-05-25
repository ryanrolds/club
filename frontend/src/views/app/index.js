import React, { useState, useEff, useEffect } from 'react'
import CssBaseline from '@material-ui/core/CssBaseline'
import Paper from '@material-ui/core/Paper'
import TopBar from '../../components/appBar/topBar'
import Seat from '../../components/atoms/seat'
import { makeStyles } from '@material-ui/core'

const useStyles = makeStyles((theme) => ({
  root: {
    display: 'flex',
    justifyContent: 'center',
    flexWrap: 'wrap',
    listStyle: 'none',
    padding: theme.spacing(0.5),
    margin: 0,
  }
}))

export default function App() {
  const classes = useStyles()
  const [seatData, setSeatData] = useState([
    { key: 0, label: 'Empty Seat' },
    { key: 1, label: 'Diego Kourchenko' },
    { key: 2, label: 'Empty Seat' },
    { key: 3, label: 'Chris Sjoblom' },
    { key: 4, label: 'Adam Stuthers' },
    { key: 5, label: 'Doctor Bud' },
    { key: 6, label: 'Ryan Olds' },
  ])

  // useEffect(() => {
  //   const seatIndex = seatData.findIndex(obj => obj.key === key)
  //   const updatedSeat = { ...seatData[seatIndex], label: 'Leavable Seat'}
  //   const updatedSeatData = [
  //     ...seatData.slice(0, seatIndex),
  //     updatedSeat,
  //     ...seatData.slice(seatIndex + 1),
  //   ]
  //   setSeatData(updatedSeatData)
  // })

  const handleJoin = (key) => {
    console.log(`JOINED SEAT: ${key}`)
  }
  const handleLeave = (event)  => {
    console.log(`LEAVE SEAT: ${event}`)
  }

  return (
    <>
      <CssBaseline />
      <main>
        <TopBar />
        {/* Paper or Main Page Component Here */}
      </main>
    </>
  )
}
