/* eslint-disable jsx-a11y/media-has-caption */
import React from 'react'
import { makeStyles } from '@material-ui/core'
import PropTypes from 'prop-types'

const useStyles = makeStyles(() => ({
  video: {
    width: '100%',
  },
}))

function PersonVideo({ person }) {
  const classes = useStyles()

  return (
    <video
      // eslint-disable-next-line no-param-reassign
      ref={(video) => { video.srcObject = person.stream }}
      autoPlay
      muted={person.muted}
      className={classes.video}
    >
      <track default />
    </video>
  )
}

PersonVideo.propTypes = {
  person: PropTypes.objectOf(PropTypes.object).isRequired,
}

export default PersonVideo
