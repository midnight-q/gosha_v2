import 'package:flutter/material.dart';

class GoshaAppBar extends StatelessWidget implements PreferredSizeWidget {
  GoshaAppBar({Key? key, required this.title, required this.onPressed, required this.name }) : super(key: key);

  final VoidCallback  onPressed;
  final String title;
  final String name;

  final preferredSize = AppBar().preferredSize;

  @override
  Widget build(BuildContext context) {
    return AppBar(
      title: Text(title),
      actions: [
        TextButton(
            onPressed: onPressed,
            child: Padding(
              padding: EdgeInsets.only(right: 20),
              child: Text(
                name.length > 0 ? name : "Click to select app",
                style: TextStyle(color: Colors.white, fontSize: 16),
              ),
            ))
      ],
    );
  }

}
