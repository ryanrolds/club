import React from 'react'
import { makeStyles, Grid } from '@material-ui/core'
import Skeleton from '@material-ui/lab/Skeleton'
import PropTypes from 'prop-types'
import PersonCard from '../person/personCard'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
}))

function GroupGridList(streamerData) {
  const classes = useStyles()
  const singer = {
    id: '123',
  }
  const local = {
    id: '9999',
  }
  const peers = [
    { id: '1' },
    { id: '2' },
    { id: '3' },
    { id: '4' },
    { id: '5' },
  ]

  return (
    <Grid container className={classes.root} spacing={2}>
      <Grid key={singer.id} item xs={4}>
        <PersonCard personId={singer.id} />
      </Grid>
      <Grid key={local.id} item xs={4}>
        <PersonCard personId={local.id} />
      </Grid>
      {peers.length ? peers.map((peer) => (
        <Grid key={peer.id} item xs={4}>
          <PersonCard personId={peer.id} />
        </Grid>
      )) : <Grid item xs={4}>
          <Skeleton />
        </Grid>}
    </Grid>
  )
}

GroupGridList.defaultProps = {
  peers: [],
}

GroupGridList.propTypes = {
  peers: PropTypes.array.isRequired
}

export default GroupGridList
