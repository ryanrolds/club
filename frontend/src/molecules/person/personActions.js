import React, { useState } from 'react'
import { makeStyles, Button, Backdrop } from '@material-ui/core'
import SpeedDial from '@material-ui/lab/SpeedDial'
import SpeedDialAction from '@material-ui/lab/SpeedDialAction'
import SpeedDialIcon from '@material-ui/lab/SpeedDialIcon'
import VolumeOffRounded from '@material-ui/icons/VolumeOffRounded'

const useStyles = makeStyles((theme) => ({
  root: {
    height: 380,
    transform: 'translateZ(0px)',
    flexGrow: 1,
  },
  speedDial: {
    position: 'absolute',
    bottom: theme.spacing(2),
    right: theme.spacing(2),
  },
}))


const actions = [
  { icon: <VolumeOffRounded />, name: 'Mute'},
]

function PersonActions() {
  const classes = useStyles()
  const [open, setOpen] = useState(false)
  const [hidden, setHidden] = useState(false)

  const handleOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div className={classes.root}>
      <Backdrop open={open} />
      <SpeedDial
        ariaLabel="Streamer Options"
        className={classes.speedDial}
        hidden={hidden}
        icon={<SpeedDialIcon />}
        onClose={handleClose}
        onOpen={handleOpen}
        open={open}
        >
          {actions.map((action) => (
            <SpeedDialAction
              key={action.name}
              icon={action.icon}
              tooltipTitle={action.name}
              tooltipOpen
              onClick={handleClose}
              />
          ))}
        </SpeedDial>
    </div>
  )
}

export default PersonActions
