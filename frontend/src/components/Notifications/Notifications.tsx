import { Notification as NotifIcon } from '@carbon/icons-react';
import { HeaderGlobalAction, ToastNotification } from '@carbon/react';
import { css } from '@emotion/css';
import { useEffect, useRef, useState } from 'react';
import { NotificationInterface } from '../../types/notifications';

interface Props {
  notifications: NotificationInterface[];
}

export const Notifications = ({ notifications }: Props): JSX.Element => {
  const [isNotificationsOpen, setIsNotificationsOpen] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setIsNotificationsOpen(false);
      }
    };
    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  return (
    <div ref={ref} className={styles.wrapper}>
      <HeaderGlobalAction aria-label="Notifications" onClick={() => setIsNotificationsOpen(!isNotificationsOpen)}>
        <NotifIcon />
      </HeaderGlobalAction>
      <div className={styles.notifs(isNotificationsOpen)}>
        {notifications.map((notif, idx) => (
          <ToastNotification key={idx} title={notif.title} subtitle={notif.subtitle} caption={notif.caption} onClose={() => false} />
        ))}
      </div>
    </div>
  );
};

const styles = {
  notifs: (isOpen: boolean) => css`
    position: absolute;
    right: 0;
    margin-top: 1rem;
    margin-right: ${isOpen ? '0rem' : '-23rem'};
    padding-right: 1rem;
    width: 23rem;
    transition: margin-right 0.3s;

    display: flex;
    flex-direction: column;
    gap: 1rem;
    align-items: end;
  `,
  wrapper: css`
    height: calc(100% - 3rem);
  `,
};
