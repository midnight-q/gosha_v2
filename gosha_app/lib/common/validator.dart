String? validateEmail(val) {
  if (val == null || val.isEmpty) {
    return "Enter email";
  }
  var isEmailValid = RegExp(
          r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+")
      .hasMatch(val);
  if (!isEmailValid) {
    return "Enter correct email";
  }
  return null;
}

String? validatePassword(val) {
  if (val == null || val.isEmpty) {
    return "Enter password";
  }
  var isEmailValid =
      RegExp(r"^(?=.*?[a-zA-Z])(?=.*?[0-9]).{8,}$").hasMatch(val);
  if (!isEmailValid) {
    return "Email must contain characters and numbers";
  }
  return null;
}
