import React from 'react'
import { Grid } from '@material-ui/core'
import Skeleton from '@material-ui/lab/Skeleton'
import PropTypes from 'prop-types'
import PersonItem from '../person/personItem'

function GroupGridList({ singer, local, peers }) {
  return (
    <Grid container spacing={2}>
      {singer ?
          <Grid key={singer.id} item xs={4}>
            <PersonItem personId={singer.id} />
          </Grid>
          :
          null
      }
      <Grid key={local.id} item xs={4}>
        <PersonItem personId={local.id} />
      </Grid>
      {peers.length ? peers.map((peer) => (
        <Grid key={peer.id} item xs={4}>
          <PersonItem personId={peer.id} />
        </Grid>
      )) : <Grid item xs={4}>
          <Skeleton />
        </Grid>}
    </Grid>
  )
}

GroupGridList.propTypes = {
  singer: PropTypes.object,
  local: PropTypes.object.isRequired,
  peers: PropTypes.array,
}

export default GroupGridList
