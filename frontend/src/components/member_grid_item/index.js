import React from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import LocalPeer from '../peer_local'
import RemotePeer from '../peer_remote'

const useStyles = makeStyles({
  gridItem: (colSize) => ({
    'grid-column': `${colSize}fr`,
  }),
})

const MemberListItem = ({
  id,
  colSize,
  name,
  localID,
  localStream,
  setLocalStream,
}) => {
  const classes = useStyles(colSize)

  return (
    <div className={classes.gridItem}>
      {localID === id && (
        <LocalPeer
          id={id}
          name={name}
          stream={localStream}
          setStream={setLocalStream}
        />
      )}
      {localID !== id && (
        <RemotePeer id={id} name={name} localStream={localStream} />
      )}
    </div>
  )
}

MemberListItem.defaultProps = {
  localStream: null,
}

MemberListItem.propTypes = {
  id: PropTypes.string.isRequired,
  colSize: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  localID: PropTypes.string.isRequired,
  localStream: PropTypes.instanceOf(MediaStream),
  setLocalStream: PropTypes.func.isRequired,
}

export default MemberListItem
