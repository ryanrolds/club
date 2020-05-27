/* eslint-disable jsx-a11y/media-has-caption */
import React from 'react';
import { makeStyles } from '@material-ui/core';
import PropTypes from 'prop-types';

const useStyles = makeStyles(() => ({
  video: {
    width: '100%',
  },
}));

function PersonVideo({ person }) {
  const classes = useStyles();
  const ref = React.createRef();
  ref.srcObject = person.stream;

  return (
    <video
      ref
      autoPlay
      muted={person.muted}
      className={classes.video}
    >
      <track default />
    </video>
  );
}

PersonVideo.propTypes = {
  person: PropTypes.objectOf(PropTypes.object).isRequired,
};

export default PersonVideo;
