import React from 'react'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import MemberList from '../../components/member_list'

const Group = ({ id, members }) => {
  return (
    <>
      <div>{id}</div>
      <MemberList members={members} />
    </>
  )
}

Group.propTypes = {
  id: PropTypes.string.isRequired,
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired
    })
  ).isRequired,
}

const mapStateToProps = (state) => {
  return {
    id: state.group.id,
    members: state.group.members,
  }
}

export default connect(mapStateToProps)(Group)
