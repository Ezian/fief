import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faEnvelope,
  faLock,
  faMailBulk,
  faSignOutAlt,
  faUser,
  faUserPlus,
} from "@fortawesome/free-solid-svg-icons";

library.add(faUser, faUserPlus, faLock, faEnvelope, faSignOutAlt);

export { FontAwesomeIcon };