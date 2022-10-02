import { useState, ChangeEvent, SetStateAction, Dispatch } from 'react';
import cl from 'classnames';

import { Status, STATUS_COLORS } from 'constants/values';
import { Status as StatusType } from 'api/proto-http/nft';

import styles from 'styles/StatusOptions.module.scss';

const OptionName = ({ optionKey }: { optionKey: any }) => (
  <div className={styles.optionName}>
    <span
      className={styles.color}
      // @ts-ignore
      style={STATUS_COLORS[optionKey]
        // @ts-ignore
        ? { backgroundColor: STATUS_COLORS[optionKey] }
        : { border: '1.5px solid white' }
      }
    />
    {optionKey}
  </div>
);

interface IStatusOptionsProps {
  activeStatus: StatusType;
  setActiveStatus: Dispatch<SetStateAction<StatusType>>;
}

export const StatusOptions = ({ activeStatus, setActiveStatus }: IStatusOptionsProps) => {
  const [isOpen, setOpenStatus] = useState(false);

  const toggleDropdown = () => setOpenStatus(v => !v);
  const handleRadioCLick = ({ target: { value } }: ChangeEvent<HTMLInputElement>) => {
    setActiveStatus(value as StatusType);
  };

  return (
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
        <OptionName optionKey={activeStatus} />
      </div>
      {isOpen &&
        <div className={styles.dropdownBody}>
          {Object.keys(Status).map((key) => (
            <div className={styles.option} key={key}>
              <OptionName optionKey={key} />
              <input
                type="radio"
                value={key}
                checked={activeStatus === key}
                onChange={handleRadioCLick}
                style={activeStatus === key ? {
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
        </div>
      }
    </div>
  )
};