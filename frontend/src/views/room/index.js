import { connect } from 'react-redux'
import GroupList from '../group_list'

const mapStateToProps = state => {
    console.log(state)

    return {
        groups: state.room.groups
    }
}

const mapDispatchToProps = dispatch => {
    return {}
}

const Room = connect(mapStateToProps, mapDispatchToProps)(GroupList)

export default Room