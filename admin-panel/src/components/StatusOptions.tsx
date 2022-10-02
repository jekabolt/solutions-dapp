import { useState, ChangeEvent, SetStateAction, Dispatch } from 'react';
import cl from 'classnames';
import OutsideClickHandler from 'react-outside-click-handler';

import { Status, STATUS_COLORS } from 'constants/values';
import { Status as StatusType } from 'api/proto-http/nft';

import styles from 'styles/StatusOptions.module.scss';

const Option = ({ optionKey, titleOption = false }: { optionKey?: StatusType, titleOption?: boolean }) => (
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

interface IStatusOptionsProps {
  activeStatus?: StatusType;
  setActiveStatus: Dispatch<SetStateAction<StatusType | undefined>>;
}

export const StatusOptions = ({ activeStatus, setActiveStatus }: IStatusOptionsProps) => {
  const [isOpen, setOpenStatus] = useState(false);
  const [selectedStatus, setSelectedStatus] = useState<StatusType | undefined>();

  const toggleDropdown = () => setOpenStatus(v => !v);
  const handleRadioCLick = ({ target: { value } }: ChangeEvent<HTMLInputElement>) => {
    setSelectedStatus(value as StatusType);
  };

  const applyStatusSelection = () => {
    setActiveStatus(selectedStatus);
    setOpenStatus(false);
  };

  const handleOutsideClick = () => {
    setSelectedStatus(activeStatus);
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
          <Option optionKey={activeStatus} titleOption />
        </div>
        {isOpen &&
          <div className={styles.dropdownBody}>
            {Object.keys(Status).map((key) => (
              <div className={styles.optionWrapper} key={key}>
                <Option optionKey={key as StatusType} />
                <input
                  type="radio"
                  value={key}
                  checked={selectedStatus === key}
                  onChange={handleRadioCLick}
                  style={selectedStatus === key ? {
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