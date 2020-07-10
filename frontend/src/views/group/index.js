import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import MemberGrid from '../../components/member_grid'

const Group = ({ localID, members }) => {
  return (
    <>
      {/* <div>{id}</div> */}
      <MemberGrid localID={localID} members={members} />
    </>
  )
}

Group.propTypes = {
  localID: PropTypes.string.isRequired,
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
    })
  ).isRequired,
}

const mapStateToProps = (state) => {
  return {
    id: state.group.id,
    members: state.group.members,
    localID: state.local.id,
  }
}

export default connect(mapStateToProps)(Group)
