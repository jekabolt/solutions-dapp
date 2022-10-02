import { Fragment, useState } from 'react';
import { useNavigate } from '@tanstack/react-location';
import { useMutation } from '@tanstack/react-query';
import cl from 'classnames';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCopy, faCaretLeft, faCaretDown } from '@fortawesome/free-solid-svg-icons'

import { Header } from 'components/Header';
import { submitTrackingNumber } from 'api';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';
import { ROUTES } from 'constants/routes';
import styles from 'styles/Nft.module.scss';

const copyToClickboard = (v: string) => navigator.clipboard.writeText(v);

// todo: add real data
const ImagesList = ({ list }: { list: any[] }) => (
  <div className={styles.imagesList}>
    {list.map((v) => (
      <Fragment key={v}>
        <div className={styles.imageBlock} />
      </Fragment>
    ))}
  </div>
);

const ShippingInfo = ({
  fullName,
  address,
  zipCode,
  city,
  country,
  email,
}: {
  fullName: string;
  address: string;
  zipCode: string;
  city: string;
  country: string;
  email: string;
  }) => {
  const [trackingNumber, setTrackingNumber] = useState('');
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY) || '';

  // add disable on loading state
  const { mutate } = useMutation(() => submitTrackingNumber(
    token,
    {
      id: 'test-id',
      trackingNumber,
    }
  ));

  const handleTrackSubmit = () => {
    // add some sort of validation
    if (trackingNumber !== '' && trackingNumber.length > 1) {
      mutate();
    }
  };

  return (
    <div className={styles.shippingInfo}>
      <div className={styles.shippingLine}>
        <span>FullName</span>
        <span>{fullName}</span>
      </div>
      <div className={styles.shippingLine}>
        <span>Address</span>
        <span>{address}</span>
      </div>
      <div className={styles.shippingLine}>
        <span>ZipCode</span>
        <span>{zipCode}</span>
      </div>
      <div className={styles.shippingLine}>
        <span>City</span>
        <span>{city}</span>
      </div>
      <div className={styles.shippingLine}>
        <span>country</span>
        <span>{country}</span>
      </div>
      <div className={styles.shippingLine}>
        <span>email</span>
        <span>{email}
          <FontAwesomeIcon onClick={() => copyToClickboard(email)} icon={faCopy} />
        </span>
      </div>
      <div className={styles.shippingLine}>
        <span>tracking number</span>
        <input
          type="text"
          onChange={({ target: { value } }) => setTrackingNumber(value)}
        />
      </div>
      <button onClick={handleTrackSubmit}>
        SUBMIT
      </button>
      {/* <div className={styles.infoLine}>
        <span>tracking number</span>
        <span className={cl(
          styles.description,
          styles.withCopy
        )}>
          {trackingNumber}
          <FontAwesomeIcon onClick={() => copyToClickboard(trackingNumber)} icon={faCopy} />
        </span>
      </div> */}
    </div>
  );
};
export const Nft = () => {
  const navigate = useNavigate();
  const [isShippingOpen, setShippingOpenStatus] = useState(false);
  const images = [1, 2, 3, 4, 5];

  const info = [
    { name: 'description', description: 'description wefwef wef' },
    { name: 'ETH address', description: '0xb794f5ea0ba39494ce839613fffba74279579268', withCopy: true },
    { name: 'mint sequence number', description: 'description wefwef wef' },
    { name: 'amount', description: 'description wefwef wef' },
    { name: 'email', description: 'long_email_liza_liza_liza_liza_liza@wefwf.com', withCopy: true },
    { name: 'burnTxid', description: '0xb794f5ea0ba39494ce839613fffba74279579268', withCopy: true },
    { name: 'burn info', description: 'description wefwef wef' },
    { name: 'offchain url', description: 'description wefwef wef' },
    { name: 'ipfs uri i.e. inchanin url', description: 'description wefwef wef' },
  ];

  const toggleShippingInfo = () => setShippingOpenStatus(v => !v);

  return (
    <div>
      <Header />
      <button onClick={() => navigate({ to: ROUTES.home })} className={styles.button}>
        <FontAwesomeIcon icon={faCaretLeft} />
        BACK TO DASHBOARD
      </button>
      <div className={styles.statusLine}>
        <span>reference images</span>
        <span className={styles.status}>
          <span className={styles.statusLabel}>status</span>
          <span className={styles.statusName}>
            <span
              className={styles.color}
            // style={{
            // backgroundColor: STATUS_COLORS[status]
            // }}
            />
            Done
          </span>
        </span>
      </div>
      <ImagesList list={images} />
      <div className={styles.mainInfo}>
        <div className={styles.info}>
          {info.map(({ name, description, withCopy }) => (
            <div key={name} className={styles.infoLine}>
              <span>{name}</span>
              <span className={cl(
                styles.description,
                withCopy && styles.withCopy
              )}>
                {description}
                {withCopy && <FontAwesomeIcon onClick={() => copyToClickboard(description)} icon={faCopy} />}
              </span>
            </div>
          ))}
          <div className={styles.shippingInfoContainer}>
            <div className={styles.infoLine} onClick={toggleShippingInfo}>
              <span>shipping information</span>
              <span className={cl(styles.arrow, isShippingOpen ? styles.arrowActive : '')}>
                <FontAwesomeIcon icon={faCaretDown} />
              </span>
            </div>
            {isShippingOpen &&
              <ShippingInfo
                fullName="Barak Obama"
                address="address"
                zipCode="01101"
                city="NY"
                country="USA"
                email="email__wdwdw_12345sc__wfwefwef@wdwd.com"  
              />
            }
          </div>
        </div>
        <div className={styles.imageBlock}>
          upload image
        </div>
      </div>
      <button className={styles.button}>
        BURN || UPLOAD IMAGE
      </button>
    </div>
  );
}
