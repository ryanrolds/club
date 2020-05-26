import React from 'react'
import { Grid } from '@material-ui/core'
import PropTypes from 'prop-types'
import PersonItem from './personItem'

function PersonGridList({ local, peers }) {
  return (
    <Grid container spacing={2}>
      <Grid key={local.id} item xs={6}>
        <PersonItem person={local} />
      </Grid>
      {peers ? peers.map((peer) => (
        <Grid key={peer.id} item xs={4}>
          <PersonItem person={peer} />
        </Grid>
      )): null}
    </Grid>
  )
}

PersonGridList.propTypes = {
  singer: PropTypes.object,
  local: PropTypes.object.isRequired,
  peers: PropTypes.array,
}

export default PersonGridList
