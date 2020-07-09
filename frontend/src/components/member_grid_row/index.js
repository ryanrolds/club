import React from 'react'
import PropTypes from 'prop-types'
import MemberGridItem from '../member_grid_item'

const MemberGridRow = ({ cols, members, localID, localStream, setLocalStream }) => {
  const colSize = cols * (cols / members.length)

  return (
    <>
      {members.map((member) => (
        <MemberGridItem
          key={member.id}
          colSize={colSize}
          id={member.id}
          name={member.name}
          localID={localID}
          localStream={localStream}
          setLocalStream={setLocalStream}
        />
      ))}
    </>
  )
}

MemberGridRow.defaultProps = {
  localStream: null,
}

MemberGridRow.propTypes = {
  cols: PropTypes.number.isRequired,
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
    }).isRequired
  ).isRequired,
  localID: PropTypes.string.isRequired,
  localStream: PropTypes.instanceOf(MediaStream),
  setLocalStream: PropTypes.func.isRequired,
}

export default MemberGridRow
