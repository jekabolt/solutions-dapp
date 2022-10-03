import { Fragment, useState, Dispatch, SetStateAction } from 'react';
import { useNavigate } from '@tanstack/react-location';
import { useMutation } from '@tanstack/react-query';
import cl from 'classnames';
import ImageUploading, { ImageListType } from 'react-images-uploading';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCopy, faCaretLeft, faCaretDown } from '@fortawesome/free-solid-svg-icons'

import { Header } from 'components/Header';
import { submitTrackingNumber, burnNft, submitNewNftMintRequest, ImageToUpload } from 'api';
import { AUTH_LOCAL_STORAGE_KEY } from 'constants/values';
import { ROUTES } from 'constants/routes';
import styles from 'styles/Nft.module.scss';

const copyToClickboard = (v: string) => navigator.clipboard.writeText(v);

const previewImages = [1, 2, 3, 4, 5];
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

export const Nft = () => {
  const navigate = useNavigate();
  const [isShippingOpen, setShippingOpenStatus] = useState(false);
  const token = localStorage.getItem(AUTH_LOCAL_STORAGE_KEY) || '';
  const [uploadedImages, setUploadedImages] = useState<ImageListType>([]);

  // add disable on loading state
  const { mutate } = useMutation(() => burnNft(
    token,
    {
      address: "string",
      id: "string",
      mintSequenceNumber: 0,
      shipping: {
        address: "string",
        city: "string",
        country: "string",
        email: "string",
        fullName: "string",
        zipCode: "string"
      },
      txid: "string"
    }
  ));

  const { mutate: uploadNftReferences } = useMutation((sampleImages: ImageToUpload[]) => submitNewNftMintRequest(
    token,
    {
      nftMintRequest: {
        TxHash: "string",
        description: "string",
        ethAddress: "string",
        id: "string",
        mintSequenceNumber: 0
      },
      sampleImages,
    }
  ));

  const toggleShippingInfo = () => setShippingOpenStatus(v => !v);

  const handleBurnAndUpload = () => {
    // to call burn request
    // mutate();

    // to call upload images request
    const imagesToUpload = getImagesToUpload(uploadedImages)
    uploadNftReferences(imagesToUpload);
  };

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
      <ImagesList list={previewImages} />
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
                token={token}
              />
            }
          </div>
        </div>
        <div className={styles.imageBlock}>
          <UploadImageComponent images={uploadedImages} setImages={setUploadedImages} />
        </div>
      </div>
      <button className={styles.button} onClick={handleBurnAndUpload}>
        BURN || UPLOAD IMAGE
      </button>
    </div>
  );
}

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
  token,
}: {
  fullName: string;
  address: string;
  zipCode: string;
  city: string;
  country: string;
  email: string;
    token: string;
  }) => {
  const [trackingNumber, setTrackingNumber] = useState('');

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

const getImagesToUpload = (imagesList: ImageListType): ImageToUpload[] =>
  imagesList.reduce((acc, image) => {
    acc.push({ raw: image.dataUrl });

    return acc;
  }, [] as ImageToUpload[]);

const UploadImageComponent = ({
  images,
  setImages,
}: {
  images: ImageListType,
  setImages: Dispatch<SetStateAction<ImageListType>>,
}) => {
  const onChange = (imageList: ImageListType) => {
    setImages(imageList);
  };

  return (
    <ImageUploading
      multiple
      value={images}
      onChange={onChange}
      maxNumber={5}
      dataURLKey="dataUrl"
    >
      {({
        imageList,
        onImageUpload,
        onImageRemoveAll,
        onImageUpdate,
        onImageRemove,
        isDragging,
        dragProps,
      }) => (
        // write your building UI
        <div className="upload__image-wrapper">
          <button
            style={isDragging ? { color: 'red' } : undefined}
            onClick={onImageUpload}
            {...dragProps}
          >
            Click or Drop here
          </button>
          &nbsp;
          <button onClick={onImageRemoveAll}>Remove all images</button>
          {imageList.map((image, index) => (
            <div key={index} className="image-item">
              <img src={image['dataUrl']} alt="" width="100" />
              <div className="image-item__btn-wrapper">
                <button onClick={() => onImageUpdate(index)}>Update</button>
                <button onClick={() => onImageRemove(index)}>Remove</button>
              </div>
            </div>
          ))}
        </div>
      )}
    </ImageUploading>
  );
}
