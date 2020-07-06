// https://github.com/facebook/react/issues/11163#issuecomment-628379291
import React, { useEffect } from 'react'
import PropTypes from 'prop-types'

const Media = ({ srcObject, ...props }) => {
  const refVideo = React.createRef()

  useEffect(() => {
    if (!refVideo.current) {
      return
    }

    refVideo.current.srcObject = srcObject
  }, [refVideo, srcObject])

  return <video ref={refVideo} {...props} /> // eslint-disable-line
}

Media.defaultProps = {
  srcObject: null,
}

Media.propTypes = {
  srcObject: PropTypes.instanceOf(MediaStream),
}

export default Media
