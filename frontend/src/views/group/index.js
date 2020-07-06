import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import MemberList from '../../components/member_list'

const Group = ({ id, localID, members }) => {
  return (
    <>
      <div>{id}</div>
      <MemberList localID={localID} members={members} />
    </>
  )
}

Group.propTypes = {
  id: PropTypes.string.isRequired,
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
