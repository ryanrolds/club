import React from 'react'
import { Grid } from '@material-ui/core'
import PropTypes from 'prop-types'
import PersonItem from './personItem'

function PersonGridList({ singer, local, peers }) {
  return (
    <Grid container spacing={2}>
      {singer.stream ?
          <Grid key={singer.id} item xs={4}>
            <PersonItem person={singer} />
          </Grid>
          :
          null
      }
      <Grid key={local.id} item xs={4}>
        <PersonItem person={local} />
      </Grid>
      {peers.map((peer) => (
        <Grid key={peer.id} item xs={4}>
          <PersonItem person={peer} />
        </Grid>
      ))}
    </Grid>
  )
}

PersonGridList.propTypes = {
  singer: PropTypes.object,
  local: PropTypes.object.isRequired,
  peers: PropTypes.array,
}

export default PersonGridList
