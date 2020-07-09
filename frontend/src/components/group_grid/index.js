import React from 'react'
import PropTypes from 'prop-types'
import { makeStyles } from '@material-ui/core/styles'
import GroupGridItem from '../group_grid_item'

const useStyles = makeStyles({
  grid: {
    display: 'grid',
    width: '100vw',
    height: '100vh',
    'grid-gap': '1rem',
    'justify-content': 'center',
    'align-content': 'center',
  },
  gridColumns: (props) => ({
    'grid-template-columns': `repeat(${props.cols}, ${100 / props.cols}vmin)`,
    'grid-template-rows': `repeat(${props.rowss}, ${100 / props.rows}vmin)`,
  }),
})

const GroupGrid = ({ groups, onGroupClick }) => {
  const classes = useStyles({ cols: 2, rows: 1 })

  return (
    <div className={`${classes.grid} ${classes.gridColumns}`}>
      {groups.map((group) => (
        <GroupGridItem
          key={group.id}
          name={group.name}
          onClick={() => onGroupClick(group.id)}
        />
      ))}
    </div>
  )
}

GroupGrid.propTypes = {
  groups: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.string.isRequired,
      name: PropTypes.string.isRequire,
      // members: PropTypes.arrayOf(PropTypes.shape({}).isRequired),
      num_members: PropTypes.number.isRequired,
    }).isRequired
  ).isRequired,
  onGroupClick: PropTypes.func.isRequired,
}

export default GroupGrid
