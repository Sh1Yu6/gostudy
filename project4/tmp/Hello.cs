using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Hello : MonoBehaviour
{
    public float MoveSpeed = 1;
    void Start()
    {

    }

    // Update is called once per frame
    void Update()
    {
       
        if (Input.GetKey(KeyCode.W))      //œÚ«∞
        {
            transform.Translate(Vector3.up * Time.deltaTime * MoveSpeed);

        }
        if (Input.GetKey(KeyCode.A))
        {
            //transform.Rotate(Vector3.up * Time.deltaTime * -MoveSpeed * 2);
            transform.Translate(Vector3.left * Time.deltaTime * MoveSpeed);
        }
        if (Input.GetKey(KeyCode.S))
        {
            transform.Translate(Vector3.down * Time.deltaTime * MoveSpeed);
        }
        if (Input.GetKey(KeyCode.D))
        {
            // transform.Rotate(Vector3.up * Time.deltaTime * MoveSpeed * 2);
            transform.Translate(Vector3.right * Time.deltaTime * MoveSpeed);
        }
       
    }
}
