import { useState, ChangeEvent, useContext } from 'react';
import cl from 'classnames';
import OutsideClickHandler from 'react-outside-click-handler';

import { Status, STATUS_COLORS } from 'constants/values';
import { Context } from 'context';
import styles from 'styles/StatusOptions.module.scss';

const Option = ({ optionKey, titleOption = false }: { optionKey?: Status, titleOption?: boolean }) => (
  <div className={cl(styles.option, titleOption ? styles.titleOption : '')}>
    <div className={styles.optionName}>
      {optionKey
        ? <span
          className={styles.color}
          style={{ backgroundColor: STATUS_COLORS[optionKey] }}
        />
        : <span />
      }
      {optionKey || "STATUSES"}
    </div>
    {titleOption && <span>&gt;</span>}
  </div>
);

export const StatusOptions = () => {
  const { state, dispatch } = useContext(Context);
  const [isOpen, setOpenStatus] = useState(false);
  const [selectedStatus, setSelectedStatus] = useState('');

  const toggleDropdown = () => setOpenStatus(v => !v);
  const handleRadioCLick = ({ target: { value } }: ChangeEvent<HTMLInputElement>) => {
    setSelectedStatus(value);
  };

  const applyStatusSelection = () => {
    dispatch({ type: 'setStatus', payload: selectedStatus as Status });
    dispatch({ type: 'setPage', payload: 1 });
    setOpenStatus(false);
  };

  const handleOutsideClick = () => {
    setSelectedStatus(state.status);
    setOpenStatus(false);
  };

  return (
    <OutsideClickHandler onOutsideClick={handleOutsideClick}>
      <div className={styles.statusOptions}>
        <div
          className={cl(
            styles.dropdownTitle,
            isOpen
              ? styles.dropdownTitleOpen
              : ''
          )}
          onClick={toggleDropdown}
        >
          <Option optionKey={state.status} titleOption />
        </div>
        {isOpen &&
          <div className={styles.dropdownBody}>
            {Object.keys(Status).map((key) => (
              <div className={styles.optionWrapper} key={key}>
                <Option optionKey={key as Status} />
                <input
                  type="radio"
                  value={key}
                  checked={(selectedStatus || state.status) === key}
                  onChange={handleRadioCLick}
                  style={(selectedStatus || state.status) === key ? {
                    // @ts-ignore
                    backgroundColor: STATUS_COLORS[key] || '#000',
                    // @ts-ignore
                    ...(!STATUS_COLORS[key] && {
                      border: "2px solid #fff",
                    }),
                  } : {}}
                />
              </div>
            ))}
            <div className={styles.dropdownBodyFooter}>
              <button onClick={applyStatusSelection} className={styles.applyButton}>APPLY</button>
            </div>
          </div>
        }
      </div>
    </OutsideClickHandler>
  )
};