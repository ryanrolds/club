import React, { useState } from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import MemberGridItem from '../member_grid_item'

const useStyles = makeStyles({
  grid: {
    display: 'grid',
    width: '100vw',
    height: '100vh',
    padding: '1rem',
    'grid-gap': '1rem',
    'justify-content': 'center',
    'align-content': 'center',
  },
  gridColumns: (props) => ({
    'grid-template-columns': `repeat(auto-fit, ${120 / props.cols}vmin)`,
  }),
})

const MemberGrid = ({ localID, members }) => {
  const [localStream, setLocalStream] = useState(null)
  const numMembers = members.length
  const sqrtMembers = Math.sqrt(numMembers)
  const numColumns = Math.ceil(sqrtMembers)
  const classes = useStyles({ cols: numColumns })
  const colSize = numColumns * (numColumns / members.length)

  return (
    <div className={`${classes.grid} ${classes.gridColumns}`}>
      <MemberGridItem
        key={localID}
        colSize={colSize}
        id={localID}
        name={localID}
        localID={localID}
        localStream={localStream}
        setLocalStream={setLocalStream}
      />
      {members.map(
        (member) =>
          member.id !== localID && (
            <MemberGridItem
              key={member.id}
              colSize={colSize}
              id={member.id}
              name={member.name}
              localID={localID}
              localStream={localStream}
              setLocalStream={setLocalStream}
            />
          )
      )}
    </div>
  )
}

MemberGrid.propTypes = {
  localID: PropTypes.string.isRequired,
  members: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequired,
    }).isRequired
  ).isRequired,
}

export default MemberGrid
