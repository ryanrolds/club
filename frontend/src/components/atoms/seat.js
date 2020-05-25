import React, { useState } from 'react'
import { makeStyles } from '@material-ui/core/styles'
import Chip from '@material-ui/core/Chip'
import FaceIcon from '@material-ui/icons/Face'
import PersonAdd from '@material-ui/icons/PersonAdd'
import CancelRounded from '@material-ui/icons/CancelRounded'

const useStyles = makeStyles((theme) => ({
  chip: {
    margin: theme.spacing(0.5),
  },
}))

export default function Seat(props) {
  const classes = useStyles()

  const handleJoinSeat = (event) => {
    props.joinFn(props.id)
  }
  const handleLeaveSeat = (event) => {
    console.log(`LEAVE SEAT: ${event}`)
  }
  const canJoin = () => {
    return props.label === 'Empty Seat'
  }

  const canLeave = () => {
    return props.label === 'Leavable Seat'
  }

  return (
    <Chip
      variant="outlined"
      color="primary"
      onDelete={canJoin() ? handleJoinSeat : handleLeaveSeat}
      deleteIcon={canJoin() ? <PersonAdd /> : canLeave() ? <CancelRounded /> : null}
      label={props.label}
      icon={<FaceIcon />}
      className={classes.chip}
    />
  )
}

Seat.defaultProps = {
  label: 'Empty Seat',
  id: '',
  joinFn: (event) => { console.log(event) },
  leaveFn: (event) => { console.log(event) }
}
